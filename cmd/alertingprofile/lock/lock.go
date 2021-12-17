package lock

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
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
				return cmderr.WrongIDArgumentFormatError
			}
			return lockRun(id)
		},
	}

	return cmd
}

func lockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.AlertingProfilesLockManagerCommand{
		ID:   id,
		Mode: types.LockedMode,
	}
	params := alerting_profiles.NewAlertingProfilesLockManagerParams().WithV(apiconfig.Version).WithBody(&body)
	_, err = apiClient.Client.AlertingProfiles.AlertingProfilesLockManager(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
