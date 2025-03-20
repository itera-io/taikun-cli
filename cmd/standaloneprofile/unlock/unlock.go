package unlock

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UnlockOptions struct {
	ID int32
}

func NewCmdUnlock() *cobra.Command {
	var opts UnlockOptions

	cmd := cobra.Command{
		Use:   "unlock <standalone-profile-id>",
		Short: "Unlock a standalone profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unlockRun(&opts)
		},
	}

	return &cmd
}

func unlockRun(opts *UnlockOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.StandAloneProfileLockManagementCommand{
		Id:   &opts.ID,
		Mode: *taikuncore.NewNullableString(&types.UnlockedMode),
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.StandaloneProfileAPI.StandaloneprofileLockManagement(context.TODO()).StandAloneProfileLockManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
