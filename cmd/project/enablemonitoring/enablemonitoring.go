package enablemonitoring

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type EnableMonitoringOptions struct {
	ProjectID int32
}

func NewCmdEnableMonitoring() *cobra.Command {
	var opts EnableMonitoringOptions

	cmd := cobra.Command{
		Use:   "enable-monitoring <project-id>",
		Short: "Enable a project's monitoring",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return enableMonitoringRun(&opts)
		},
	}

	return &cmd
}

func enableMonitoringRun(opts *EnableMonitoringOptions) (err error) {
	isMonitoringEnabled, err := getMonitoringState(opts.ProjectID)
	if err != nil {
		return
	}
	if isMonitoringEnabled {
		err = cmderr.ErrProjectMonitoringAlreadyEnabled
		return
	}
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

func getMonitoringState(projectID int32) (isMonitoringEnabled bool, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersDetailsParams().WithV(taikungoclient.Version)
	params = params.WithProjectID(projectID)

	response, err := apiClient.Client.Servers.ServersDetails(params, apiClient)
	if err != nil {
		return
	}
	isMonitoringEnabled = response.Payload.Project.IsMonitoringEnabled

	return
}
