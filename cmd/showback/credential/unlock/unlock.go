package unlock

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := cobra.Command{
		Use:   "unlock <showback-credential-id",
		Short: "Unlock a showback credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			showbackCredentialID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return unlockRun(showbackCredentialID)
		},
	}

	return &cmd
}

func unlockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.ShowbackCredentialLockCommand{ID: id, Mode: types.UnlockedMode}
	params := showback.NewShowbackLockManagerParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Showback.ShowbackLockManager(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
