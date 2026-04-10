package delete

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
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
			return deleteAccount(accountID)
		},
	}

	return &cmd
}

func deleteAccount(accountID int32) (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.AccountsAPI.AccountsDelete(context.TODO(), accountID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("account", accountID)
	return nil
}
