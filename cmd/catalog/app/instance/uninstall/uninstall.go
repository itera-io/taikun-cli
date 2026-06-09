package uninstall

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type UninstallOptions struct {
	ProjectAppId int32
}

func NewCmdUninstall() *cobra.Command {
	var opts UninstallOptions

	cmd := cobra.Command{
		Use:   "uninstall <APP_INSTANCE_ID>",
		Short: "Trigger uninstall for application instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectAppId, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return uninstallAppRun(cmd, opts)
		},
	}

	return &cmd
}

func uninstallAppRun(cmd *cobra.Command, opts UninstallOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	_, response, err := myApiClient.Client.ProjectAppsAPI.ProjectappDelete(ctx, opts.ProjectAppId).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("application instance", opts.ProjectAppId)

	return nil
}
