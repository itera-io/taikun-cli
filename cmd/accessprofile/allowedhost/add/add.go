package add

import (
	"fmt"
	"net"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/allowed_host"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"IP-ADDRESS", "ip-address",
		),
		field.NewVisible(
			"MASK-BITS", "mask-bits",
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
				return fmt.Errorf("Mask bits must be in the range of [0, 32]")
			}

			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.IpAddress, "ip-address", "ip", "", "Ip Address (required)")
	cmdutils.MarkFlagRequired(cmd, "ip-address")

	cmd.Flags().Int32VarP(&opts.MaskBits, "mask-bits", "m", 0, "Mask bits (required)")
	cmdutils.MarkFlagRequired(cmd, "mask-bits")

	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "Description")

	cmdutils.AddOutputOnlyIDFlag(cmd)
	cmdutils.AddColumnsFlag(cmd, addFields)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.CreateAllowedHostCommand{
		AccessProfileID: opts.AccessProfileID,
		IPAddress:       opts.IpAddress,
		MaskBits:        opts.MaskBits,
		Description:     opts.Description,
	}

	params := allowed_host.NewAllowedHostCreateParams().WithV(taikungoclient.Version).WithBody(&body)

	response, err := apiClient.Client.AllowedHost.AllowedHostCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
