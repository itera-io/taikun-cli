package add

import (
	"context"
	"fmt"
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
	Url            string
	Username       string
	Password       string
	Datacenter     string
	ResourcePool   string
	DataStore      string
	DrsEnabled     bool
	Hypervisors    []string
	VmTemplate     string
	Continent      string
	OrganizationID int32

	PrivateNetworkName string
	PrivateNetwork     string
	PrivateNetmask     int32
	PrivateGateway     string
	PrivateBeginRange  string
	PrivateEndRange    string

	PublicNetworkName string
	PublicNetwork     string
	PublicNetmask     int32
	PublicGateway     string
	PublicBeginRange  string
	PublicEndRange    string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add an vSphere cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if ((opts.DrsEnabled) && (len(opts.Hypervisors) != 0)) || (!opts.DrsEnabled) && (len(opts.Hypervisors) == 0) {
				return fmt.Errorf("Specify only one of [--drs-enabled,--hypervisors]")
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVar(&opts.Url, "url", "", "vSphere API host (required)")
	cmdutils.MarkFlagRequired(&cmd, "url")

	cmd.Flags().StringVar(&opts.Username, "username", "", "vSphere username (required)")
	cmdutils.MarkFlagRequired(&cmd, "username")

	cmd.Flags().StringVar(&opts.Password, "password", "", "vSphere password (required)")
	cmdutils.MarkFlagRequired(&cmd, "password")

	cmd.Flags().StringVar(&opts.Datacenter, "datacenter", "", "vSphere storage (required)")
	cmdutils.MarkFlagRequired(&cmd, "datacenter")

	cmd.Flags().StringVar(&opts.ResourcePool, "resource-pool", "", "vSphere VM template (required)")
	cmdutils.MarkFlagRequired(&cmd, "resource-pool")

	cmd.Flags().StringVar(&opts.DataStore, "data-store", "", "vSphere data store (required)")
	cmdutils.MarkFlagRequired(&cmd, "data-store")

	cmd.Flags().BoolVar(&opts.DrsEnabled, "drs-enabled", false, "vSphere DRS enabled (required)")

	cmd.Flags().StringSliceVar(&opts.Hypervisors, "hypervisors", []string{}, "vSphere hypervisors in format: \"hypervisor1,hypervisor2,...\" (required)")

	cmd.Flags().StringVar(&opts.VmTemplate, "vm-template", "", "Vm template for vSphere (required)")
	cmdutils.MarkFlagRequired(&cmd, "vm-template")

	cmd.Flags().StringVar(&opts.Continent, "continent", "Europe", "vSphere continent (optional)")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization", "o", 0, "vSphere organization ID (optional)")

	// Private network
	cmd.Flags().StringVar(&opts.PrivateNetworkName, "private-network-name", "", "vSphere private network name (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-network-name")

	cmd.Flags().StringVar(&opts.PrivateNetwork, "private-network", "", "vSphere private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-network")

	cmd.Flags().Int32Var(&opts.PrivateNetmask, "private-netmask", 0, "vSphere private netmask (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-netmask")

	cmd.Flags().StringVar(&opts.PrivateGateway, "private-gateway", "", "vSphere private gateway (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-gateway")

	cmd.Flags().StringVar(&opts.PrivateBeginRange, "private-begin-range", "", "vSphere begin of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-begin-range")

	cmd.Flags().StringVar(&opts.PrivateEndRange, "private-end-range", "", "vSphere end of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "private-end-range")

	// Public network
	cmd.Flags().StringVar(&opts.PrivateNetworkName, "public-network-name", "", "vSphere private network name (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-network-name")

	cmd.Flags().StringVar(&opts.PublicNetwork, "public-network", "", "vSphere private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-network")

	cmd.Flags().Int32Var(&opts.PublicNetmask, "public-netmask", 0, "vSphere private netmask (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-netmask")

	cmd.Flags().StringVar(&opts.PublicGateway, "public-gateway", "", "vSphere private gateway (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-gateway")

	cmd.Flags().StringVar(&opts.PublicBeginRange, "public-begin-range", "", "vSphere begin of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-begin-range")

	cmd.Flags().StringVar(&opts.PublicEndRange, "public-end-range", "", "vSphere end of the range of the private network (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-end-range")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Get Datacenter ID
	datacenterBody := taikuncore.DatacenterListCommand{
		Url:            *taikuncore.NewNullableString(&opts.Url),
		Username:       *taikuncore.NewNullableString(&opts.Username),
		Password:       *taikuncore.NewNullableString(&opts.Password),
		DatacenterName: *taikuncore.NewNullableString(&opts.Datacenter),
	}
	dataDC, responseDC, errDC := myApiClient.Client.VsphereCloudCredentialAPI.VsphereDatacenterList(context.TODO()).DatacenterListCommand(datacenterBody).Execute()
	if errDC != nil {
		err = tk.CreateError(responseDC, errDC)
		return
	}
	if len(dataDC) != 1 {
		return fmt.Errorf("Datacenter ID query had multiple responses.")
	}
	datacenterID := dataDC[0].GetDatacenter()

	// Prepare the arguments for the query
	publicNetwork := taikuncore.CreateVsphereNetworkDto{
		Name:                 *taikuncore.NewNullableString(&opts.PublicNetworkName),
		Gateway:              *taikuncore.NewNullableString(&opts.PublicGateway),
		IpAddress:            *taikuncore.NewNullableString(&opts.PublicNetwork),
		NetMask:              &opts.PublicNetmask,
		BeginAllocationRange: *taikuncore.NewNullableString(&opts.PublicBeginRange),
		EndAllocationRange:   *taikuncore.NewNullableString(&opts.PublicEndRange),
	}
	privateNetwork := taikuncore.CreateVsphereNetworkDto{
		Name:                 *taikuncore.NewNullableString(&opts.PrivateNetworkName),
		Gateway:              *taikuncore.NewNullableString(&opts.PrivateGateway),
		IpAddress:            *taikuncore.NewNullableString(&opts.PrivateNetwork),
		NetMask:              &opts.PrivateNetmask,
		BeginAllocationRange: *taikuncore.NewNullableString(&opts.PrivateBeginRange),
		EndAllocationRange:   *taikuncore.NewNullableString(&opts.PrivateEndRange),
	}
	body := taikuncore.CreateVsphereCommand{
		Name:             *taikuncore.NewNullableString(&opts.Name),
		Username:         *taikuncore.NewNullableString(&opts.Username),
		Url:              *taikuncore.NewNullableString(&opts.Url),
		Password:         *taikuncore.NewNullableString(&opts.Password),
		DatacenterName:   *taikuncore.NewNullableString(&opts.Datacenter),
		DatacenterId:     *taikuncore.NewNullableString(&datacenterID),
		DatastoreName:    *taikuncore.NewNullableString(&opts.DataStore),
		ResourcePoolName: *taikuncore.NewNullableString(&opts.ResourcePool),
		DrsEnabled:       &opts.DrsEnabled,
		VmTemplateName:   *taikuncore.NewNullableString(&opts.VmTemplate),
		Continent:        *taikuncore.NewNullableString(&opts.Continent),
		OrganizationId:   *taikuncore.NewNullableInt32(&opts.OrganizationID),
		Hypervisors:      opts.Hypervisors,
		PublicNetwork:    &publicNetwork,
		PrivateNetwork:   &privateNetwork,
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.VsphereCloudCredentialAPI.VsphereCreate(context.TODO()).CreateVsphereCommand(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	return out.PrintResult(data, addFields)

}
