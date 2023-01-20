package disable

import (
	"errors"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/autoscaling"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
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
	_, err = isAutoscalingEnabled(opts.ProjectID)
	if err != nil {
		return
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.DisableAutoscalingCommand{
		ProjectID: opts.ProjectID,
	}

	params := autoscaling.NewAutoscalingDisableAutoscalingParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Autoscaling.AutoscalingDisableAutoscaling(params, apiClient)
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
			err = errors.New("Project autoscaling already disabled.")
		}
	}
	return
}
