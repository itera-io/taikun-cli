package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/flavors"
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
		field.NewHidden(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
	},
)

type AddOptions struct {
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
		Short: "Add a server to a project",
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.ServerForCreateDto{
		DiskSize:  types.GiBToB(opts.DiskSize),
		Flavor:    opts.Flavor,
		Name:      opts.Name,
		ProjectID: opts.ProjectID,
		Role:      types.GetServerRole(opts.Role),
	}

	if len(opts.KubernetesNodeLabels) != 0 {
		body.KubernetesNodeLabels, err = parseKubernetesNodeLabelsFlag(opts.KubernetesNodeLabels)
		if err != nil {
			return
		}
	}

	params := servers.NewServersCreateParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Servers.ServersCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload, addFields)
	}

	return
}

func parseKubernetesNodeLabelsFlag(labelsData []string) ([]*models.KubernetesNodeLabelsDto, error) {
	labels := make([]*models.KubernetesNodeLabelsDto, len(labelsData))
	for i, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, errors.New("Invalid empty kubernetes node label")
		}
		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid kubernetes node label format: %s", labelData)
		}
		labels[i] = &models.KubernetesNodeLabelsDto{
			Key:   tokens[0],
			Value: tokens[1],
		}
	}
	return labels, nil
}

func flavorCompletionFunc(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
	completions = make([]string, 0)

	if len(args) == 0 {
		return
	}
	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return
	}

	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := flavors.NewFlavorsGetSelectedFlavorsForProjectParams().WithV(api.Version)
	params = params.WithProjectID(&projectID)

	for {
		response, err := apiClient.Client.Flavors.FlavorsGetSelectedFlavorsForProject(params, apiClient)
		if err != nil {
			return
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

	return
}
