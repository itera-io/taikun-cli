package edit

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type EditOptions struct {
	ProjectID int32
	MinSize   int32
	MaxSize   int32
}

func NewCmdEdit() *cobra.Command {
	var opts EditOptions

	cmd := cobra.Command{
		Use:   "edit <project-id>",
		Short: "Edit autoscaling for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return editRun(&opts)
		},
	}

	cmd.Flags().Int32Var(&opts.MaxSize, "max-size", 1, "The autoscaler's maximum size")
	cmd.Flags().Int32Var(&opts.MinSize, "min-size", 0, "The autoscaler's minimum size")

	return &cmd
}

func editRun(opts *EditOptions) (err error) {
	autoscalingEnabled, err := cmdutils.IsAutoscalingEnabled(opts.ProjectID)
	if err != nil {
		return err
	}
	if !autoscalingEnabled {
		err = fmt.Errorf("project autoscaling is disabled and thus cannot be edited")
		return err
	}

	myApiClient := tk.NewClient()
	body := taikuncore.EditAutoscalingCommand{
		ProjectId: &opts.ProjectID,
		MinSize:   &opts.MinSize,
		MaxSize:   &opts.MaxSize,
	}
	response, err := myApiClient.Client.AutoscalingAPI.AutoscalingEdit(context.TODO()).EditAutoscalingCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
