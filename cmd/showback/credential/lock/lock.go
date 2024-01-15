package lock

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(showbackCredentialID)
		},
	}

	return &cmd
}

func lockRun(showbackCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikunshowback.ShowbackCredentialLockCommand{
		Id:   &showbackCredentialID,
		Mode: *taikunshowback.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.ShowbackClient.ShowbackCredentialsAPI.ShowbackcredentialsLockManagement(context.TODO()).ShowbackCredentialLockCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
