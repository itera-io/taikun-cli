package disable2fa

import (
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return disable2fa(cmd, &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.VerificationCode, "code", "c", "", "Verification code")

	return &cmd
}

func disable2fa(cmd *cobra.Command, opts *Disable2faOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.NewDisableTwoFaManagementCommand()
	if opts.VerificationCode != "" {
		body.SetVerificationCode(opts.VerificationCode)
	}

	response, err := myApiClient.Client.AccountsAPI.AccountsDisable2faManagement(ctx).DisableTwoFaManagementCommand(*body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
