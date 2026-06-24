package syncpackage

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type SyncOptions struct {
	ProjectAppId int32
}

func NewCmdSync() *cobra.Command {
	var opts SyncOptions

	cmd := cobra.Command{
		Use:   "sync <APP_INSTANCE_ID>",
		Short: "Trigger sync for application instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectAppId, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return syncAppRun(cmd, opts)
		},
	}

	return &cmd
}

func syncAppRun(cmd *cobra.Command, opts SyncOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.SyncProjectAppCommand{ProjectAppId: &opts.ProjectAppId}

	response, err := myApiClient.Client.ProjectAppsAPI.ProjectappSync(ctx).SyncProjectAppCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil
}
