package list

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"IP-ADDRESS", "ipAddress",
		),
		field.NewVisible(
			"ACCESS-PROFILE", "accessProfileName",
		),
		field.NewVisible(
			"MASK-BITS", "maskBits",
		),
		field.NewVisible(
			"DESCRIPTION", "description",
		),
	},
)

type ListOptions struct {
	AccessProfileID int32
	Limit           int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <access-profile-id>",
		Short: "List an access profile's allowed hosts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accessProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.AccessProfileID = accessProfileID
			return listRun(cmd, &opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(cmd, listFields)

	return cmd
}

func listRun(cmd *cobra.Command, opts *ListOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.AllowedHostAPI.AllowedhostList(ctx, opts.AccessProfileID)

	var allowedHosts []interface{}
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		for _, host := range data.GetData() {
			allowedHosts = append(allowedHosts, host)
		}

		count := int32(len(allowedHosts))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}
		if count == data.GetTotalCount() {
			break
		}
		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(allowedHosts)) > opts.Limit {
		allowedHosts = allowedHosts[:opts.Limit]
	}

	return out.PrintResults(allowedHosts, listFields)
}
