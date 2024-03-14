package add

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
			"ORG", "organizationName",
		),
		field.NewVisible(
			"CONTINENT", "continentName",
		),
		field.NewVisible(
			"URL", "url",
		),
		field.NewHidden(
			"STORAGE", "storage",
		),
		field.NewHidden(
			"VM-TEMPLATE-NAME", "vmTemplateName",
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	Name           string
	ApiHost        string
	ClientId       string
	ClientSecret   string
	Storage        string
	VmTemplate     string
	Hypervisors    []string
	OrganizationID int32
	Continent      string

	PrivateNetwork    string
	PrivateNetmask    int32
	PrivateGateway    string
	PrivateBeginRange string
	PrivateEndRange   string
	PrivateBridge     string

	PublicNetwork    string
	PublicNetmask    int32
	PublicGateway    string
	PublicBeginRange string
	PublicEndRange   string
	PublicBridge     string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add an Proxmox cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVar(&opts.ApiHost, "api-host", "", "Proxmox API host (required)")
	cmdutils.MarkFlagRequired(&cmd, "api-host")

	cmd.Flags().StringVar(&opts.ClientId, "client-id", "", "Proxmox client ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "client-id")

	cmd.Flags().StringVar(&opts.ClientSecret, "client-secret", "", "Proxmox client secret (required)")
	cmdutils.MarkFlagRequired(&cmd, "client-secret")

	cmd.Flags().StringVar(&opts.Storage, "storage", "", "Proxmox storage (required)")
	cmdutils.MarkFlagRequired(&cmd, "storage")

	cmd.Flags().StringVar(&opts.VmTemplate, "vm-template", "", "Proxmox VM template (required)")
	cmdutils.MarkFlagRequired(&cmd, "vm-template")

	cmd.Flags().StringSliceVar(&opts.Hypervisors, "hypervisors", []string{}, "Proxmox hypervisors in format: \"hypervisor1,hypervisor2,...\" (required)")
	cmdutils.MarkFlagRequired(&cmd, "hypervisors")

	cmd.Flags().StringVar(&opts.Continent, "continent", "", "Proxmox continent (optional)")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization", "o", 0, "Proxmox organization ID (optional)")

	// Private network
	cmd.Flags().StringVar(&opts.PrivateNetwork, "private-network", "", "Proxmox private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-network")

	cmd.Flags().Int32Var(&opts.PrivateNetmask, "private-netmask", 0, "Proxmox private netmask (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-netmask")

	cmd.Flags().StringVar(&opts.PrivateGateway, "private-gateway", "", "Proxmox private gateway (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-gateway")

	cmd.Flags().StringVar(&opts.PrivateBeginRange, "private-begin-range", "", "Proxmox begin of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-begin-range")

	cmd.Flags().StringVar(&opts.PrivateEndRange, "private-end-range", "", "Proxmox end of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-end-range")

	cmd.Flags().StringVar(&opts.PrivateBridge, "private-bridge", "", "Proxmox client ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-bridge")

	// Public network
	cmd.Flags().StringVar(&opts.PublicNetwork, "public-network", "", "Proxmox private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-network")

	cmd.Flags().Int32Var(&opts.PublicNetmask, "public-netmask", 0, "Proxmox private netmask (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-netmask")

	cmd.Flags().StringVar(&opts.PublicGateway, "public-gateway", "", "Proxmox private gateway (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-gateway")

	cmd.Flags().StringVar(&opts.PublicBeginRange, "public-begin-range", "", "Proxmox begin of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-begin-range")

	cmd.Flags().StringVar(&opts.PublicEndRange, "public-end-range", "", "Proxmox end of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-end-range")

	cmd.Flags().StringVar(&opts.PublicBridge, "public-bridge", "", "Proxmox client ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-bridge")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	publicNetwork := taikuncore.CreateProxmoxNetworkDto{
		Bridge:               *taikuncore.NewNullableString(&opts.PublicBridge),
		Gateway:              *taikuncore.NewNullableString(&opts.PublicGateway),
		IpAddress:            *taikuncore.NewNullableString(&opts.PublicNetwork),
		NetMask:              &opts.PublicNetmask,
		BeginAllocationRange: *taikuncore.NewNullableString(&opts.PublicBeginRange),
		EndAllocationRange:   *taikuncore.NewNullableString(&opts.PublicEndRange),
	}
	privateNetwork := taikuncore.CreateProxmoxNetworkDto{
		Bridge:               *taikuncore.NewNullableString(&opts.PrivateBridge),
		Gateway:              *taikuncore.NewNullableString(&opts.PrivateGateway),
		IpAddress:            *taikuncore.NewNullableString(&opts.PrivateNetwork),
		NetMask:              &opts.PrivateNetmask,
		BeginAllocationRange: *taikuncore.NewNullableString(&opts.PrivateBeginRange),
		EndAllocationRange:   *taikuncore.NewNullableString(&opts.PrivateEndRange),
	}
	body := taikuncore.CreateProxmoxCommand{
		Name:           *taikuncore.NewNullableString(&opts.Name),
		TokenId:        *taikuncore.NewNullableString(&opts.ClientId),
		Url:            *taikuncore.NewNullableString(&opts.ApiHost),
		TokenSecret:    *taikuncore.NewNullableString(&opts.ClientSecret),
		Storage:        *taikuncore.NewNullableString(&opts.Storage),
		VmTemplateName: *taikuncore.NewNullableString(&opts.VmTemplate),
		Continent:      *taikuncore.NewNullableString(&opts.Continent),
		OrganizationId: *taikuncore.NewNullableInt32(&opts.OrganizationID),
		Hypervisors:    opts.Hypervisors,
		PublicNetwork:  &publicNetwork,
		PrivateNetwork: &privateNetwork,
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.ProxmoxCloudCredentialAPI.ProxmoxCreate(context.TODO()).CreateProxmoxCommand(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	return out.PrintResult(data, addFields)

}
