package create

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Email     string
	CreateOrg bool
}

func NewCmdCreateAccount() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <ACCOUNT_NAME>",
		Short: "Create a new account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createAccount(args[0], &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email")
	cmd.Flags().BoolVarP(&opts.CreateOrg, "create-org", "", false, "Create organization")
	_ = cmd.MarkFlagRequired("email")

	return &cmd
}

func createAccount(accountName string, opts *CreateOptions) (err error) {
	// input parameters sanity check
	if opts.Email == "" {
		return fmt.Errorf("email must be specified")
	}
	myApiClient := tk.NewClient()

	body := taikuncore.CreateAccountCommand{
		Name:               *taikuncore.NewNullableString(&accountName),
		Email:              *taikuncore.NewNullableString(&opts.Email),
		CreateOrganization: &opts.CreateOrg,
	}

	_, response, err := myApiClient.Client.AccountsAPI.AccountsCreate(context.TODO()).CreateAccountCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
