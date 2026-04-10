package available

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
	},
)

func NewCmdListAvailable() *cobra.Command {
	cmd := cobra.Command{
		Use:   "available <ACCOUNT_ID>",
		Short: "List available organizations in an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			accountID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return listAvailableRun(accountID)
		},
	}

	return &cmd
}

func listAvailableRun(accountID int32) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.AccountsAPI.AccountsAccountOrganizationsAvailable(context.TODO(), accountID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResults(data, listFields)
}
