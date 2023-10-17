package disablemonitoring

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	isMonitoringEnabled, err := cmdutils.IsMonitoringEnabled(opts.ProjectID)
	if err != nil {
		return
	}
	if !isMonitoringEnabled {
		err = cmderr.ErrProjectMonitoringAlreadyDisabled
		return
	}
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	body := taikuncore.MonitoringOperationsCommand{
		ProjectId: &opts.ProjectID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.ProjectsAPI.ProjectsMonitoring(context.TODO()).MonitoringOperationsCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	out.PrintStandardSuccess()
	return

	/*
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
	*/
}
