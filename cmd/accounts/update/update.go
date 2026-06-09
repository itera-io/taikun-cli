package update

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type UpdateOptions struct {
	AccountID int32
	Name      string
	Email     string
}

func NewCmdUpdate() *cobra.Command {
	var opts UpdateOptions

	cmd := cobra.Command{
		Use:   "update <ACCOUNT_ID>",
		Short: "Update account",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AccountID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return updateRun(cmd, &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email")

	return &cmd
}

func updateRun(cmd *cobra.Command, opts *UpdateOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.UpdateAccountCommand{
		Id: &opts.AccountID,
	}

	if opts.Name != "" {
		body.SetName(opts.Name)
	}
	if opts.Email != "" {
		body.SetEmail(opts.Email)
	}

	_, response, err := myApiClient.Client.AccountsAPI.AccountsUpdate(ctx).UpdateAccountCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
