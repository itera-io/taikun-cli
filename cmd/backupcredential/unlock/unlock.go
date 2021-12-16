package unlock

import (
	"taikun-cli/api"
	"taikun-cli/utils"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <backup-credential-id>",
		Short: "Unlock a backup credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			backupCredentialID, err := types.Atoi32(args[0])
			if err != nil {
				return utils.WrongIDArgumentFormatError
			}
			return unlockRun(backupCredentialID)
		},
	}

	return cmd
}

func unlockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.BackupLockManagerCommand{
		ID:   id,
		Mode: types.UnlockedMode,
	}
	params := s3_credentials.NewS3CredentialsLockManagerParams().WithV(utils.ApiVersion).WithBody(&body)
	_, err = apiClient.Client.S3Credentials.S3CredentialsLockManager(params, apiClient)
	if err == nil {
		utils.PrintStandardSuccess()
	}

	return
}
