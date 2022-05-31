package list

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/security_group"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"PROFILE", "profileName",
		),
		field.NewVisible(
			"REMOTE-IP-PREFIX", "remoteIpPrefix",
		),
		field.NewVisibleWithToStringFunc(
			"PROTOCOL", "protocol", out.FormatStringUpper,
		),
		field.NewVisible(
			"MIN-PORT", "portMinRange",
		),
		field.NewVisible(
			"MAX-PORT", "portMaxRange",
		),
	},
)

type ListOptions struct {
	StandAloneProfileID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <standalone-profile-id>",
		Short: "List a standalone profile's security groups",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandAloneProfileID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := security_group.NewSecurityGroupListParams().WithV(taikungoclient.Version)
	params = params.WithStandAloneProfileID(opts.StandAloneProfileID)

	response, err := apiClient.Client.SecurityGroup.SecurityGroupList(params, apiClient)
	if err == nil {
		return out.PrintResults(response.Payload, listFields)
	}

	return
}
