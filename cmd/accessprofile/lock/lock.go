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
		Use:   "lock <access-profile-id>",
		Short: "Lock an access profile",
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

func lockRun(accessProfileID int32) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.AccessProfilesLockManagementCommand{
		Id:   &accessProfileID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}
	response, err := myApiClient.Client.AccessProfilesAPI.AccessprofilesLockManager(context.TODO()).AccessProfilesLockManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
