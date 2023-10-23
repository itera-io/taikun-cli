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

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <policy-profile-id>",
		Short: "Unlock a policy profile",
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

func unlockRun(policyProfileID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.OpaProfileLockManagerCommand{
		Id:   &policyProfileID,
		Mode: *taikuncore.NewNullableString(&types.UnlockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.OpaProfilesAPI.OpaprofilesLockManager(context.TODO()).OpaProfileLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

	/*
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
	*/
}
