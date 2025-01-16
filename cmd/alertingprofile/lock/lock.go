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
		Use:   "lock <alerting-profile-id>",
		Short: "Lock an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(id)
		},
	}

	return cmd
}

func lockRun(alertingProfileID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.AlertingProfilesLockManagerCommand{
		Id:   &alertingProfileID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesLockManager(context.TODO()).AlertingProfilesLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
