package cancel

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CancelOptions struct {
	ProjectAppId int32
}

func NewCmdCancel() *cobra.Command {
	var opts CancelOptions

	cmd := cobra.Command{
		Use:   "cancel <APP_INSTANCE_ID>",
		Short: "Trigger cancel for application instance",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectAppId, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return cancelAppRun(opts)
		},
	}

	return &cmd
}

func cancelAppRun(opts CancelOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.CancelProjectAppCommand{
		ProjectAppId: &opts.ProjectAppId,
	}

	response, err := myApiClient.Client.ProjectAppsAPI.ProjectappCancel(context.TODO()).CancelProjectAppCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil
}
