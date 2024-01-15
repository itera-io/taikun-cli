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
		Use:   "lock <kubernetes-profile-id>",
		Short: "Lock a kubernetes profile",
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

func lockRun(kubernetesProfileID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.KubernetesProfilesLockManagerCommand{
		Id:   &kubernetesProfileID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.KubernetesProfilesAPI.KubernetesprofilesLockManager(context.TODO()).KubernetesProfilesLockManagerCommand(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	out.PrintStandardSuccess()
	return

}
