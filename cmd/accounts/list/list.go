package list

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

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("ORGS", "organizationsCount"),
		field.NewVisible("USERS", "usersCount"),
		field.NewVisible("GROUPS", "groupsCount"),
		field.NewVisible("PROJECTS", "projectsCount"),
	},
)

type ListOptions struct {
	Limit  int32
	Search string
}

func NewCmdListAccounts() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List available accounts",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listAccounts(&opts)
		},
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmd.Flags().StringVarP(&opts.Search, "search", "s", "", "Search string")
	cmdutils.AddColumnsFlag(&cmd, listFields)
	return &cmd
}

func listAccounts(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	var accounts = make([]taikuncore.AccountList, 0)

	req := myApiClient.Client.AccountsAPI.AccountsListAccounts(context.TODO())
	if opts.Search != "" {
		req = req.Search(opts.Search)
	}

	for {
		data, response, err := req.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}
		accounts = append(accounts, data.GetData()...)

		accountsCount := int32(len(accounts))
		if opts.Limit != 0 && accountsCount >= opts.Limit {
			break
		}

		nextCursor, ok := data.GetNextCursorOk()
		if !data.HasMore || !ok || nextCursor == nil {
			break
		}

		req = req.CursorId(*nextCursor)
	}

	if opts.Limit != 0 && int32(len(accounts)) > opts.Limit {
		accounts = accounts[:opts.Limit]
	}

	return out.PrintResults(accounts, listFields)
}
