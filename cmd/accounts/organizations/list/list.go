package list

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("ORG-ID", "orgId"),
		field.NewVisible("ORG-NAME", "orgName"),
		field.NewVisible("BOUND", "isBound"),
		field.NewVisible("ACCESS-LEVEL", "accessLevel"),
	},
)

type ListOptions struct {
	AccountID int32
	Limit     int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <ACCOUNT_ID>",
		Short: "List organizations in an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccountID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return listRun(&opts)
		},
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	var organizations = make([]taikuncore.OrganizationsWithGroupInfoResultDto, 0)

	req := myApiClient.Client.AccountsAPI.AccountsAccountOrganizationsWithGroup(context.TODO(), opts.AccountID).Limit(opts.Limit)
	for {
		data, response, err := req.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		organizations = append(organizations, data.GetData()...)

		organizationsCount := int32(len(organizations))
		if opts.Limit != 0 && organizationsCount >= opts.Limit {
			break
		}

		nextCursor, ok := data.GetNextCursorOk()
		if !data.HasMore || !ok || nextCursor == nil {
			break
		}

		req = req.CursorId(*nextCursor)
	}

	if opts.Limit != 0 && int32(len(organizations)) > opts.Limit {
		organizations = organizations[:opts.Limit]
	}

	return out.PrintResults(organizations, listFields)
}
