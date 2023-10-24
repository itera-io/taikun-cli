package lock

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <backup-credential-id>",
		Short: "Lock a backup credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			backupCredentialID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(backupCredentialID)
		},
	}

	return cmd
}

func lockRun(backupCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.BackupLockManagerCommand{
		Id:   &backupCredentialID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.S3CredentialsAPI.S3credentialsLockManagement(context.TODO()).BackupLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
