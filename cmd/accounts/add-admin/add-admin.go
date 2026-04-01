package addadmin

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type AddAdminOptions struct {
	AccountID int32
	UserID    string
}

func NewCmdAddAdmin() *cobra.Command {
	var opts AddAdminOptions

	cmd := cobra.Command{
		Use:   "add-admin <ACCOUNT_ID>",
		Short: "Add account admin",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccountID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return addAdminRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.UserID, "user-id", "u", "", "User ID")
	_ = cmd.MarkFlagRequired("user-id")

	return &cmd
}

func addAdminRun(opts *AddAdminOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.AddAccountAdminCommand{
		AccountId: *taikuncore.NewNullableInt32(&opts.AccountID),
		UserId:    *taikuncore.NewNullableString(&opts.UserID),
	}

	response, err := myApiClient.Client.AccountsAPI.AccountsAddAccountAdmin(context.TODO()).AddAccountAdminCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
