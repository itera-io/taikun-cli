package unlock

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <access-profile-id>",
		Short: "Unlock an access profile",
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

func unlockRun(accessProfileID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.AccessProfilesLockManagementCommand{
		ID:   accessProfileID,
		Mode: types.UnlockedMode,
	}
	params := access_profiles.NewAccessProfilesLockManagerParams().WithV(taikungoclient.Version).WithBody(&body)

	_, err = apiClient.Client.AccessProfiles.AccessProfilesLockManager(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
