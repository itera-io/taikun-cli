package uninstall

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
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
			return uninstallAppRun(opts)
		},
	}

	return &cmd
}

func uninstallAppRun(opts UninstallOptions) (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.ProjectAppsAPI.ProjectappDelete(context.TODO(), opts.ProjectAppId).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("application instance", opts.ProjectAppId)

	return nil
}
