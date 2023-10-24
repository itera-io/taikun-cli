package enablemonitoring

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
	isMonitoringEnabled, err := cmdutils.IsMonitoringEnabled(opts.ProjectID)
	if err != nil {
		return
	}
	if isMonitoringEnabled {
		err = cmderr.ErrProjectMonitoringAlreadyEnabled
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

}
