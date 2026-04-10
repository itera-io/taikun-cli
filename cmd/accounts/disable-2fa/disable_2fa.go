package disable2fa

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type Disable2faOptions struct {
	VerificationCode string
}

func NewCmdDisable2fa() *cobra.Command {
	var opts Disable2faOptions

	cmd := cobra.Command{
		Use:   "disable-2fa",
		Short: "Disable 2FA management",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return disable2fa(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.VerificationCode, "code", "c", "", "Verification code")

	return &cmd
}

func disable2fa(opts *Disable2faOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.NewDisableTwoFaManagementCommand()
	if opts.VerificationCode != "" {
		body.SetVerificationCode(opts.VerificationCode)
	}

	response, err := myApiClient.Client.AccountsAPI.AccountsDisable2faManagement(context.TODO()).DisableTwoFaManagementCommand(*body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
