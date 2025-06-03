package add

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"net"

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
			"IP-ADDRESS", "ipAddress",
		),
		field.NewVisible(
			"MASK-BITS", "maskBits",
		),
		field.NewVisible(
			"DESCRIPTION", "description",
		),
	},
)

type AddOptions struct {
	AccessProfileID int32
	IpAddress       string
	Description     string
	MaskBits        int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <access-profile-id>",
		Short: "Add an allowed host",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccessProfileID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}

			if net.ParseIP(opts.IpAddress) == nil {
				return fmt.Errorf("IP address must be valid")
			}

			if opts.MaskBits < 0 && opts.MaskBits > 32 {
				return fmt.Errorf("mask bits must be in the range of [0, 32]")
			}

			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.IpAddress, "ip-address", "i", "", "Ip Address (required)")
	cmdutils.MarkFlagRequired(cmd, "ip-address")

	cmd.Flags().Int32VarP(&opts.MaskBits, "mask-bits", "m", 0, "Mask bits (required)")
	cmdutils.MarkFlagRequired(cmd, "mask-bits")

	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "Description")

	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CreateAllowedHostCommand{
		AccessProfileId: &opts.AccessProfileID,
		Description:     *taikuncore.NewNullableString(&opts.Description),
		IpAddress:       *taikuncore.NewNullableString(&opts.IpAddress),
		MaskBits:        &opts.MaskBits,
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.AllowedHostAPI.AllowedhostCreate(context.TODO()).CreateAllowedHostCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return out.PrintResult(data, addFields)

}
