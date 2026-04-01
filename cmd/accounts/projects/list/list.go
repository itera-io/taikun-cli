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
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
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
		Short: "List projects in an account",
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
	var projects = make([]taikuncore.CommonDropdownDto, 0)

	req := myApiClient.Client.AccountsAPI.AccountsAccountProjectsDropdown(context.TODO(), opts.AccountID)

	for {
		data, response, err := req.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		projects = append(projects, data.GetData()...)

		projectsCount := int32(len(projects))
		if opts.Limit != 0 && projectsCount >= opts.Limit {
			break
		}

		nextCursor, ok := data.GetNextCursorOk()
		if !data.HasMore || !ok || nextCursor == nil {
			break
		}

		req = req.CursorId(*nextCursor)
	}

	if opts.Limit != 0 && int32(len(projects)) > opts.Limit {
		projects = projects[:opts.Limit]
	}

	return out.PrintResults(projects, listFields)
}
