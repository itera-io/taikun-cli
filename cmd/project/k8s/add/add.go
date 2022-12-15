package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"CPU", "cpu",
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "ram", out.FormatBToGiB,
		),
		field.NewVisibleWithToStringFunc(
			"DISK", "diskSize", out.FormatBToGiB,
		),
		field.NewVisible(
			"ROLE", "role",
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewHiddenWithToStringFunc(
			"CLOUDTYPE", "cloudType", out.FormatCloudType,
		),
		field.NewVisible(
			"AVAILABILITY-ZONE", "availabilityZone",
		),
		field.NewHidden(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
	},
)

type AddOptions struct {
	AvailabilityZone     string
	DiskSize             int
	Flavor               string
	KubernetesNodeLabels []string
	Name                 string
	ProjectID            int32
	Role                 string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <project-id>",
		Short: "Add a Kubernetes server to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if err := cmdutils.CheckFlagValue("role", opts.Role, types.ServerRoles); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AvailabilityZone, "availability-zone", "a", "", "Availability zone (for AWS, GCP and Azure projects)")
	cmdutils.SetFlagCompletionFunc(&cmd, "availability-zone", availabilityZoneCompletionFunc)

	cmd.Flags().IntVarP(&opts.DiskSize, "disk-size", "d", 30, "Disk size in GB")
	cmd.Flags().StringSliceVarP(&opts.KubernetesNodeLabels, "kubernetes-node-labels", "k", []string{}, "Kubernetes node labels (format: \"key=value,key2=value2,...\")")

	cmd.Flags().StringVarP(&opts.Flavor, "flavor", "f", "", "Flavor (required)")
	cmdutils.MarkFlagRequired(&cmd, "flavor")
	cmdutils.SetFlagCompletionFunc(&cmd, "flavor", flavorCompletionFunc)

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "name")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	cmdutils.MarkFlagRequired(&cmd, "role")
	cmdutils.SetFlagCompletionValues(&cmd, "role", types.ServerRoles.Keys()...)

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.ServerForCreateDto{
		AvailabilityZone: opts.AvailabilityZone,
		DiskSize:         types.GiBToB(opts.DiskSize),
		Flavor:           opts.Flavor,
		Name:             opts.Name,
		ProjectID:        opts.ProjectID,
		Role:             types.GetServerRole(opts.Role),
	}

	if len(opts.KubernetesNodeLabels) != 0 {
		body.KubernetesNodeLabels, err = parseKubernetesNodeLabelsFlag(opts.KubernetesNodeLabels)
		if err != nil {
			return
		}
	}

	params := servers.NewServersCreateParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Servers.ServersCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}

func parseKubernetesNodeLabelsFlag(labelsData []string) ([]*models.KubernetesNodeLabelsDto, error) {
	labels := make([]*models.KubernetesNodeLabelsDto, len(labelsData))

	for labelIndex, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, errors.New("Invalid empty kubernetes node label")
		}

		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid kubernetes node label format: %s", labelData)
		}

		labels[labelIndex] = &models.KubernetesNodeLabelsDto{
			Key:   tokens[0],
			Value: tokens[1],
		}
	}

	return labels, nil
}

func flavorCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	if len(args) == 0 {
		return []string{}
	}

	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return []string{}
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return []string{}
	}

	params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(taikungoclient.Version)
	params = params.WithProjectID(&projectID)

	completions := make([]string, 0)

	for {
		response, err := apiClient.Client.Flavors.FlavorsGetSelectedFlavorsForProject(params, apiClient)
		if err != nil {
			return []string{}
		}

		for _, flavor := range response.Payload.Data {
			completions = append(completions, flavor.Name)
		}

		count := int32(len(completions))

		if count == response.Payload.TotalCount {
			break
		}

		params = params.WithOffset(&count)
	}

	return completions
}

func availabilityZoneCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {

	if len(args) == 0 {
		return []string{}
	}

	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return []string{}
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return []string{}
	}

	projectParams := projects.NewProjectsListParams().WithV(taikungoclient.Version)
	projectParams = projectParams.WithID(&projectID)
	projectResponse, err := apiClient.Client.Projects.ProjectsList(projectParams, apiClient)
	if err != nil || len(projectResponse.GetPayload().Data) != 1 {
		return []string{}
	}

	projectOrgId := projectResponse.GetPayload().Data[0].OrganizationID
	ccType := projectResponse.GetPayload().Data[0].CloudType
	ccName := projectResponse.GetPayload().Data[0].CloudCredentialName
	if ccType == "OPENSTACK" {
		return []string{}
	}

	completions := make([]string, 0)

	ccParams := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(taikungoclient.Version).WithOrganizationID(&projectOrgId)
	ccResponse, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(ccParams, apiClient)
	countCC := ccResponse.GetPayload().TotalCountOpenstack + ccResponse.GetPayload().TotalCountAws + ccResponse.GetPayload().TotalCountAzure + ccResponse.GetPayload().TotalCountGoogle
	if err != nil || countCC == 0 {
		return []string{}
	}

	if ccType == "AWS" {
		amazonCCs := ccResponse.GetPayload().Amazon
		for i := 0; i < int(ccResponse.Payload.TotalCountGoogle); i++ {
			if ccName == amazonCCs[i].Name {
				completions = append(completions, amazonCCs[i].AvailabilityZones...)
			}
		}
	}
	if ccType == "AZURE" {
		azureCCs := ccResponse.GetPayload().Azure
		for i := 0; i < int(ccResponse.Payload.TotalCountAzure); i++ {
			if ccName == azureCCs[i].Name {
				completions = append(completions, azureCCs[i].AvailabilityZones...)
			}
		}
	}
	if ccType == "GOOGLE" {
		googleCCs := ccResponse.GetPayload().Google
		for i := 0; i < int(ccResponse.Payload.TotalCountGoogle); i++ {
			if ccName == googleCCs[i].Name {
				completions = append(completions, googleCCs[i].Zones...)
			}
		}
	}

	return completions
}
