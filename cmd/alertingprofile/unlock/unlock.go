package unlock

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <alerting-profile-id>",
		Short: "Unlock an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unlockRun(cmd, id)
		},
	}

	return cmd
}

func unlockRun(cmd *cobra.Command, alertingProfileID int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.AlertingProfilesLockManagerCommand{
		Id:   &alertingProfileID,
		Mode: *taikuncore.NewNullableString(&types.UnlockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesLockManager(ctx).AlertingProfilesLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
