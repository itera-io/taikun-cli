package unlock

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <cloud-credential-id>",
		Short: "Unlock a cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unlockRun(id)
		},
	}

	return cmd
}

func unlockRun(cloudCredentialID int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CloudLockManagerCommand{
		ID:   cloudCredentialID,
		Mode: types.UnlockedMode,
	}
	params := cloud_credentials.NewCloudCredentialsLockManagerParams().WithV(api.Version).WithBody(body)

	_, err = apiClient.Client.CloudCredentials.CloudCredentialsLockManager(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
