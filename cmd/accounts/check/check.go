package check

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func NewCmdCheck() *cobra.Command {
	cmd := cobra.Command{
		Use:   "check <ACCOUNT_NAME>",
		Short: "Check if an account name is available",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkAccount(args[0])
		},
	}

	return &cmd
}

func checkAccount(accountName string) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.CheckDuplicateAccountCommand{
		Name: *taikuncore.NewNullableString(&accountName),
	}

	response, err := myApiClient.Client.AccountsAPI.AccountsCheckDuplicateEntity(context.TODO()).CheckDuplicateAccountCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
