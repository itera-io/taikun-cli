package add

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			"SIZE", "size",
		),
		field.NewVisible(
			"VOLUME-TYPE", "volumeType",
		),
		field.NewVisible(
			"DEVICE-NAME", "deviceName",
		),
		field.NewVisible(
			"LUN-ID", "lunId",
		),
		field.NewVisible(
			"STATUS", "status",
		),
	},
)

type AddOptions struct {
	Name                string
	OpenStackVolumeType string
	Size                int64
	StandaloneVMID      int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <vm-id>",
		Short: "Add a disk to a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name (required)")
	cmdutils.MarkFlagRequired(&cmd, "name")

	cmd.Flags().StringVarP(&opts.OpenStackVolumeType, "openstack-volume-type", "v", "", "Volume type (for OpenStack only)")

	cmd.Flags().Int64VarP(&opts.Size, "size", "s", 0, "Size in GiB (required)")
	cmdutils.MarkFlagRequired(&cmd, "size")

	cmdutils.AddColumnsFlag(&cmd, addFields)
	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CreateStandAloneDiskCommand{
		StandaloneVmId: &opts.StandaloneVMID,
		Name:           *taikuncore.NewNullableString(&opts.Name),
		Size:           &opts.Size,
	}
	if opts.OpenStackVolumeType != "" {
		body.SetVolumeType(opts.OpenStackVolumeType)
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.StandaloneVMDisksAPI.StandalonevmdisksCreate(context.TODO()).CreateStandAloneDiskCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResult(data, addFields)

}
