package check

import (
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdCheck() *cobra.Command {
	cmd := cobra.Command{
		Use:   "check <ACCOUNT_NAME>",
		Short: "Check if an account name is available",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkAccount(cmd, args[0])
		},
	}

	return &cmd
}

func checkAccount(cmd *cobra.Command, accountName string) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.CheckDuplicateAccountCommand{
		Name: *taikuncore.NewNullableString(&accountName),
	}

	response, err := myApiClient.Client.AccountsAPI.AccountsCheckDuplicateEntity(ctx).CheckDuplicateAccountCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
