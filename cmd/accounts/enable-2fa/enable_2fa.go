package enable2fa

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdEnable2fa() *cobra.Command {
	cmd := cobra.Command{
		Use:   "enable-2fa",
		Short: "Enable 2FA management",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return enable2fa()
		},
	}

	return &cmd
}

func enable2fa() (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.AccountsAPI.AccountsEnable2faManagement(context.TODO()).Body(map[string]interface{}{}).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
