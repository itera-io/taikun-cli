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

type LockOptions struct {
	ID int32
}

func NewCmdLock() *cobra.Command {
	var opts LockOptions

	cmd := cobra.Command{
		Use:   "lock <standalone-profile-id>",
		Short: "Lock a standalone profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(&opts)
		},
	}

	return &cmd
}

func lockRun(opts *LockOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.StandAloneProfileLockManagementCommand{
		Id:   &opts.ID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneProfileAPI.StandaloneprofileLockManagement(context.TODO()).StandAloneProfileLockManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
