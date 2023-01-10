package edit

import (
	"errors"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/autoscaling"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
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
	cmd.Flags().Int32Var(&opts.MinSize, "min-size", 1, "The autoscaler's minimum size")

	return &cmd
}

func editRun(opts *EditOptions) (err error) {
	_, err = isAutoscalingEnabled(opts.ProjectID)
	if err != nil {
		return
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.EditAutoscalingCommand{
		ProjectID: opts.ProjectID,
		MaxSize:   opts.MaxSize,
		MinSize:   opts.MinSize,
	}

	params := autoscaling.NewAutoscalingEditAutoscalingParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Autoscaling.AutoscalingEditAutoscaling(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}
	return
}

func isAutoscalingEnabled(projectID int32) (res bool, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(taikungoclient.Version)
	params = params.WithProjectID(projectID)

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err == nil {
		res := response.Payload.Project.IsAutoscalingEnabled
		if !res {
			err = errors.New("Project autoscaling is disabled.\nPlease use the command 'taikun project autoscaler enable ...' instead.")
		}
	}
	return
}
