package unlock

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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
			return unlockRun(id)
		},
	}

	return cmd
}

func unlockRun(alertingProfileID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.AlertingProfilesLockManagerCommand{
		ID:   alertingProfileID,
		Mode: types.LockedMode,
	}
	params := alerting_profiles.NewAlertingProfilesLockManagerParams().WithV(taikungoclient.Version).WithBody(&body)

	_, err = apiClient.Client.AlertingProfiles.AlertingProfilesLockManager(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
