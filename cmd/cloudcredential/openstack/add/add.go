package add

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
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
			"NAME", "cloudCredentialName",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	Name               string
	Username           string
	Password           string
	URL                string
	Project            string
	Domain             string
	Region             string
	PublicNetwork      string
	AvailabilityZone   string
	InternalSubnetId   string
	VolumeType         string
	ImportNetwork      bool
	OrganizationID     int32
	OpenStackContinent string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add an OpenStack cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if opts.OpenStackContinent != "" {
				if err := cmdutils.CheckFlagValue("continent", opts.OpenStackContinent, types.OpenstackContinent); err != nil {
					return err
				}
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "OpenStack Username (required)")
	cmdutils.MarkFlagRequired(&cmd, "username")

	cmd.Flags().StringVarP(&opts.Password, "password", "p", "", "OpenStack Password (required)")
	cmdutils.MarkFlagRequired(&cmd, "password")

	cmd.Flags().StringVarP(&opts.Domain, "domain", "d", "", "OpenStack Domain (required)")
	cmdutils.MarkFlagRequired(&cmd, "domain")

	cmd.Flags().StringVar(&opts.URL, "url", "", "OpenStack URL (required)")
	cmdutils.MarkFlagRequired(&cmd, "url")

	cmd.Flags().StringVar(&opts.Project, "project", "", "OpenStack Project (required)")
	cmdutils.MarkFlagRequired(&cmd, "project")

	cmd.Flags().StringVarP(&opts.Region, "region", "r", "", "OpenStack Region (required)")
	cmdutils.MarkFlagRequired(&cmd, "region")

	cmd.Flags().StringVarP(&opts.OpenStackContinent, "continent", "c", "", "OpenStack Continent [eu,us,as] (required)")
	cmdutils.SetFlagCompletionValues(&cmd, "continent", types.OpenstackContinent.Keys()...)

	cmd.Flags().StringVar(&opts.PublicNetwork, "public-network", "", "OpenStack Public Network (required)")
	cmdutils.MarkFlagRequired(&cmd, "public-network")

	cmd.Flags().StringVar(&opts.AvailabilityZone, "availability-zone", "", "OpenStack Availability Zone")
	cmd.Flags().StringVar(&opts.InternalSubnetId, "internal-subnet-id", "", "OpenStack Internal Subnet ID")
	cmd.Flags().StringVar(&opts.VolumeType, "volume-type", "", "OpenStack Volume Type")
	cmd.Flags().BoolVar(&opts.ImportNetwork, "import-network", false, "Import Network (false by default)")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	importNetwork := opts.ImportNetwork
	body := taikuncore.CreateOpenstackCloudCommand{
		Name:                      *taikuncore.NewNullableString(&opts.Name),
		OpenStackUser:             *taikuncore.NewNullableString(&opts.Username),
		OpenStackPassword:         *taikuncore.NewNullableString(&opts.Password),
		OpenStackUrl:              *taikuncore.NewNullableString(&opts.URL),
		OpenStackProject:          *taikuncore.NewNullableString(&opts.Project),
		OpenStackPublicNetwork:    *taikuncore.NewNullableString(&opts.PublicNetwork),
		OpenStackAvailabilityZone: *taikuncore.NewNullableString(&opts.AvailabilityZone),
		OpenStackDomain:           *taikuncore.NewNullableString(&opts.Domain),
		OpenStackRegion:           *taikuncore.NewNullableString(&opts.Region),
		OpenStackVolumeType:       *taikuncore.NewNullableString(&opts.VolumeType),
		OpenStackImportNetwork:    &importNetwork,
		OpenStackInternalSubnetId: *taikuncore.NewNullableString(&opts.InternalSubnetId),
		OrganizationId:            *taikuncore.NewNullableInt32(&opts.OrganizationID),
	}
	openstackContinent := types.GetOpenstackContinent(opts.OpenStackContinent)
	if openstackContinent != nil {
		body.SetOpenStackContinent(openstackContinent.(string))
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.OpenstackCloudCredentialAPI.OpenstackCreate(context.TODO()).CreateOpenstackCloudCommand(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	return out.PrintResult(data, addFields)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := &models.CreateOpenstackCloudCommand{
			Name:                      opts.Name,
			OpenStackAvailabilityZone: opts.AvailabilityZone,
			OpenStackDomain:           opts.Domain,
			OpenStackImportNetwork:    opts.ImportNetwork,
			OpenStackInternalSubnetID: opts.InternalSubnetId,
			OpenStackPassword:         opts.Password,
			OpenStackProject:          opts.Project,
			OpenStackPublicNetwork:    opts.PublicNetwork,
			OpenStackRegion:           opts.Region,
			OpenStackURL:              opts.URL,
			OpenStackUser:             opts.Username,
			OpenStackVolumeType:       opts.VolumeType,
			OrganizationID:            opts.OrganizationID,
		}

		params := openstack.NewOpenstackCreateParams().WithV(taikungoclient.Version).WithBody(body)

		response, err := apiClient.Client.Openstack.OpenstackCreate(params, apiClient)
		if err == nil {
			return out.PrintResult(response.Payload, addFields)
		}

		return
	*/
}
