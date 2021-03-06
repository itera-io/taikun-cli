package lock

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/opa_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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
			return lockRun(id)
		},
	}

	return cmd
}

func lockRun(policyProfileID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.OpaProfileLockManagerCommand{
		ID:   policyProfileID,
		Mode: types.LockedMode,
	}
	params := opa_profiles.NewOpaProfilesLockManagerParams().WithV(taikungoclient.Version).WithBody(&body)

	_, err = apiClient.Client.OpaProfiles.OpaProfilesLockManager(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
