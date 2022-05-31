package monitoringtoggle

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type MonitoringToggleOptions struct {
	ProjectID int32
}

func NewCmdMonitoringToggle() *cobra.Command {
	var opts MonitoringToggleOptions

	cmd := cobra.Command{
		Use:   "monitoring-toggle <project-id>",
		Short: "Toggle a project's monitoring on or off",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return monitoringToggleRun(&opts)
		},
		Aliases: []string{"monitoring"},
	}

	return &cmd
}

func monitoringToggleRun(opts *MonitoringToggleOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.MonitoringOperationsCommand{ProjectID: opts.ProjectID}
	params := projects.NewProjectsMonitoringOperationsParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Projects.ProjectsMonitoringOperations(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
