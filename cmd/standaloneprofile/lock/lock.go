package lock

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone_profile"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type LockOptions struct {
	ID int32
}

func NewCmdLock() *cobra.Command {
	var opts LockOptions

	cmd := cobra.Command{
		Use:   "lock <standalone-profile-id>",
		Short: "Lock a standalone profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return lockRun(&opts)
		},
	}

	return &cmd
}

func lockRun(opts *LockOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.StandAloneProfileLockManagementCommand{
		ID:   opts.ID,
		Mode: types.LockedMode,
	}

	params := stand_alone_profile.NewStandAloneProfileLockManagementParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneProfile.StandAloneProfileLockManagement(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
