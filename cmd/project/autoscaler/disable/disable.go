package disable

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DisableOptions struct {
	ProjectID int32
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <project-id>",
		Short: "Disable the project's autoscaling",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return disableRun(&opts)
		},
	}

	return &cmd
}

func disableRun(opts *DisableOptions) (err error) {
	autoscalingEnabled, err := cmdutils.IsAutoscalingEnabled(opts.ProjectID)
	if err != nil {
		return err
	}
	if !autoscalingEnabled {
		err = fmt.Errorf("Project autoscaling already disabled")
		return err
	}

	myApiClient := tk.NewClient()
	body := taikuncore.DisableAutoscalingCommand{
		ProjectId: &opts.ProjectID,
	}
	_, response, err := myApiClient.Client.AutoscalingAPI.AutoscalingDisable(context.TODO()).DisableAutoscalingCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
