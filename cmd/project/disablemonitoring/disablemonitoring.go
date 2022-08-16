package disablemonitoring

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

type DisableMonitoringOptions struct {
	ProjectID int32
}

func NewCmdDisableMonitoring() *cobra.Command {
	var opts DisableMonitoringOptions

	cmd := cobra.Command{
		Use:   "disable-monitoring <project-id>",
		Short: "Disable a project's monitoring",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return disableMonitoringRun(&opts)
		},
	}

	return &cmd
}

func disableMonitoringRun(opts *DisableMonitoringOptions) (err error) {
	isMonitoringEnabled, err := getMonitoringState(opts.ProjectID)
	if err != nil {
		return
	}
	if !isMonitoringEnabled {
		err = cmderr.ErrProjectMonitoringAlreadyDisabled
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
