package info

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var infoFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("CLAIM-NAME", "claimName"),
		field.NewVisible("CREATED-AT", "createdAt"),
		field.NewVisible("ORGANIZATIONS", "organizations"),
		field.NewVisible("USERS", "users"),
	},
)

type InfoOptions struct {
	AccountID int32
	GroupID   int32
}

func NewCmdInfo() *cobra.Command {
	var opts InfoOptions

	cmd := cobra.Command{
		Use:   "info <ACCOUNT_ID> <GROUP_ID>",
		Short: "Get detailed information about a group in an account",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccountID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.GroupID, err = types.Atoi32(args[1])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return infoRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, infoFields)

	return &cmd
}

func infoRun(opts *InfoOptions) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.AccountsAPI.AccountsAccountGroupDetails(context.TODO(), opts.AccountID, opts.GroupID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, infoFields)
}
