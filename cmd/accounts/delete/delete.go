package delete

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdDeleteAccount() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <ACCOUNT_ID>",
		Short: "Delete an account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return deleteAccount(cmd, accountID)
		},
	}

	return &cmd
}

func deleteAccount(cmd *cobra.Command, accountID int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.AccountsAPI.AccountsDelete(ctx, accountID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("account", accountID)
	return nil
}
