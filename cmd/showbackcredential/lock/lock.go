package lock

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

func NewCmdLock() *cobra.Command {
	cmd := cobra.Command{
		Use:   "lock <showback-credential-id",
		Short: "Lock a showback credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			showbackCredentialID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			return lockRun(showbackCredentialID)
		},
	}

	return &cmd
}

func lockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.ShowbackCredentialLockCommand{ID: id, Mode: types.LockedMode}
	params := showback.NewShowbackLockManagerParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Showback.ShowbackLockManager(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
