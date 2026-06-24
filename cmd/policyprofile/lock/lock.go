package lock

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <policy-profile-id>",
		Short: "Lock a policy profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(cmd, id)
		},
	}

	return cmd
}

func lockRun(cmd *cobra.Command, policyProfileID int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.OpaProfileLockManagerCommand{
		Id:   &policyProfileID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.OpaProfilesAPI.OpaprofilesLockManager(ctx).OpaProfileLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
