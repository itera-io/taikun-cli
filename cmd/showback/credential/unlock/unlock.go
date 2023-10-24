package unlock

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return unlockRun(showbackCredentialID)
		},
	}

	return &cmd
}

func unlockRun(showbackCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikunshowback.ShowbackCredentialLockCommand{
		Id:   &showbackCredentialID,
		Mode: *taikunshowback.NewNullableString(&types.UnlockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.ShowbackClient.ShowbackCredentialsAPI.ShowbackcredentialsLockManagement(context.TODO()).ShowbackCredentialLockCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.ShowbackCredentialLockCommand{ID: showbackCredentialID, Mode: types.UnlockedMode}
		params := showback_credentials.NewShowbackCredentialsLockManagerParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.ShowbackClient.ShowbackCredentials.ShowbackCredentialsLockManager(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
