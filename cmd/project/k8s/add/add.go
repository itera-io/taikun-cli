package add

import (
	"context"
	"errors"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
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
		field.NewHiddenWithToStringFunc(
			"AVAILABILITY-ZONE", "availabilityZone", out.FormatAvailabilityZones,
		),
		field.NewHidden(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
		field.NewVisible(
			"WASM", "wasmEnabled",
		),
		field.NewHidden(
			"HYPERVISOR", "hypervisor",
		),
		field.NewHidden(
			"PROXMOX-ROLE", "proxmoxRole",
		),
		field.NewHidden(
			"PROXMOX-DISK", "proxmoxExtraDiskSize",
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
	WasmEnabled          bool
	ProxmoxDisk          int32
	Hypervisor           string
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

	cmd.Flags().StringVarP(&opts.AvailabilityZone, "availability-zone", "a", "", "Availability zone (only for AWS, GCP and Azure projects)")
	cmdutils.SetFlagCompletionFunc(&cmd, "availability-zone", availabilityZoneCompletionFunc)

	cmd.Flags().IntVarP(&opts.DiskSize, "disk-size", "d", 30, "Disk size in GB")
	cmd.Flags().StringSliceVarP(&opts.KubernetesNodeLabels, "kubernetes-node-labels", "k", []string{}, "Kubernetes node labels (format: \"key=value,key2=value2,...\")")

	cmd.Flags().StringVarP(&opts.Flavor, "flavor", "f", "", "Flavor (required)")
	cmdutils.MarkFlagRequired(&cmd, "flavor")
	cmdutils.SetFlagCompletionFunc(&cmd, "flavor", cmdutils.FlavorCompletionFunc)

	cmd.Flags().BoolVar(&opts.WasmEnabled, "enable-wasm", false, "Enable support for WASM")
	cmd.Flags().StringVar(&opts.Hypervisor, "hypervisor", "", "Proxmox hypervisor")
	cmd.Flags().Int32Var(&opts.ProxmoxDisk, "proxmox-disk", 0, "Proxmox extra disk size (this will automatically enable proxmox NFS type)")

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "name")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role [Kubemaster, Kubeworker, Bastion] (required)")
	cmdutils.MarkFlagRequired(&cmd, "role")
	cmdutils.SetFlagCompletionValues(&cmd, "role", types.ServerRoles.Keys()...)

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()
	diskSizeValue := types.GiBToB(opts.DiskSize)
	serverRole := types.GetServerRole(opts.Role)
	body := taikuncore.ServerForCreateDto{
		AvailabilityZone: *taikuncore.NewNullableString(&opts.AvailabilityZone),
		DiskSize:         &diskSizeValue,
		Flavor:           *taikuncore.NewNullableString(&opts.Flavor),
		Name:             *taikuncore.NewNullableString(&opts.Name),
		ProjectId:        &opts.ProjectID,
		Role:             &serverRole,
		WasmEnabled:      &opts.WasmEnabled,
	}

	// Only set if optional Proxmox parameter is present
	if opts.ProxmoxDisk != 0 {
		proxmoxRole, err1 := getProxmoxRole(opts.ProjectID)
		if err1 != nil {
			return err1
		}
		body.SetProxmoxRole(*proxmoxRole)
		body.SetProxmoxExtraDiskSize(opts.ProxmoxDisk)
	}

	// Only set if optional Proxmox parameter is present
	if opts.Hypervisor != "" {
		body.SetHypervisor(opts.Hypervisor)
	}

	if len(opts.KubernetesNodeLabels) != 0 {
		body.KubernetesNodeLabels, err = parseKubernetesNodeLabelsFlag(opts.KubernetesNodeLabels)
		if err != nil {
			return
		}
	}
	data, response, err := myApiClient.Client.ServersAPI.ServersCreate(context.TODO()).ServerForCreateDto(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)

}

func parseKubernetesNodeLabelsFlag(labelsData []string) ([]taikuncore.KubernetesNodeLabelsDto, error) {
	labels := make([]taikuncore.KubernetesNodeLabelsDto, len(labelsData))

	for labelIndex, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, errors.New("Invalid empty kubernetes node label")
		}

		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid kubernetes node label format: %s", labelData)
		}

		labels[labelIndex] = taikuncore.KubernetesNodeLabelsDto{
			Key:   *taikuncore.NewNullableString(&tokens[0]),
			Value: *taikuncore.NewNullableString(&tokens[1]),
		}
	}

	return labels, nil

}

func availabilityZoneCompletionFunc(cmd *cobra.Command, args []string, toComplete string) []string {
	if len(args) == 0 {
		return []string{}
	}

	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return []string{}
	}

	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ProjectsAPI.ProjectsList(context.TODO()).Id(projectID).Execute()
	if err != nil || len(data.GetData()) != 1 {
		fmt.Println(fmt.Errorf(tk.CreateError(response, err).Error()))
		return []string{}
	}

	projectOrgId := data.Data[0].GetOrganizationId()
	ccType := data.Data[0].GetCloudType()
	ccName := data.Data[0].GetCloudCredentialName()
	if ccType == "OPENSTACK" {
		return []string{}
	}

	completions := make([]string, 0)

	dataCC, responseCC, err := myApiClient.Client.CloudCredentialAPI.CloudcredentialsDashboardList(context.TODO()).OrganizationId(projectOrgId).Execute()
	if err != nil {
		fmt.Println(fmt.Errorf(tk.CreateError(responseCC, err).Error()))
		return []string{}
	}
	countCC := dataCC.GetTotalCountOpenstack() + dataCC.GetTotalCountAws() + dataCC.GetTotalCountAzure() + dataCC.GetTotalCountGoogle()

	if err != nil || countCC == 0 {
		return []string{}
	}

	if ccType == "AWS" {
		amazonCCs := dataCC.GetAmazon()
		for i := 0; i < int(dataCC.GetTotalCountAws()); i++ {
			if ccName == amazonCCs[i].GetName() {
				completions = append(completions, amazonCCs[i].AvailabilityZones...)
			}
		}
	}
	if ccType == "AZURE" {
		azureCCs := dataCC.GetAzure()
		for i := 0; i < int(dataCC.GetTotalCountAzure()); i++ {
			if ccName == azureCCs[i].GetName() {
				completions = append(completions, azureCCs[i].AvailabilityZones...)
			}
		}
	}
	if ccType == "GOOGLE" {
		googleCCs := dataCC.GetGoogle()
		for i := 0; i < int(dataCC.GetTotalCountGoogle()); i++ {
			if ccName == googleCCs[i].GetName() {
				completions = append(completions, googleCCs[i].Zones...)
			}
		}
	}

	return completions

}

// Proxmox Role type for a k8s server depends on Proxmox type specified in the Kubernetes profile.
// The names in profile don't match the names we send to the server, for it, we have this function.
func getProxmoxRole(projectId int32) (returnRole *taikuncore.ProxmoxRole, returnErr error) {
	myclient := tk.NewClient()
	data, response, err := myclient.Client.ServersAPI.ServersDetails(context.TODO(), projectId).Execute()
	if err != nil {
		returnErr = tk.CreateError(response, err)
		return
	}
	kubernetesProfile := data.GetProject()
	var proxmoxRoleString string
	switch kubernetesProfile.GetProxmoxStorage() {
	case "NFS":
		proxmoxRoleString = "NFS"
	case "OpenEBS":
		proxmoxRoleString = "STORAGE"
	default:
		proxmoxRoleString = ""
	}

	returnRole, returnErr = taikuncore.NewProxmoxRoleFromValue(proxmoxRoleString)
	return
}
