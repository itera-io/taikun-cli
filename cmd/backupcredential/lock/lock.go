package lock

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <backup-credential-id>",
		Short: "Lock a backup credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			backupCredentialID, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given ID must be a number")
			}
			return lockRun(backupCredentialID)
		},
	}

	return cmd
}

func lockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.BackupLockManagerCommand{
		ID:   id,
		Mode: cmdutils.LockedMode,
	}
	params := s3_credentials.NewS3CredentialsLockManagerParams().WithV(cmdutils.ApiVersion).WithBody(&body)
	_, err = apiClient.Client.S3Credentials.S3CredentialsLockManager(params, apiClient)
	if err == nil {
		fmt.Println("Backup credential locked")
	}

	return
}