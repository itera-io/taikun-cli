package disablemonitoring

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()
	body := taikuncore.DeploymentDisableMonitoringCommand{
		ProjectId: &opts.ProjectID,
	}

	// Execute a query into the API + graceful exit
	_, err = myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDisableMonitoring(context.TODO()).DeploymentDisableMonitoringCommand(body).Execute()
	if err != nil {
		return cmderr.ErrProjectBackupAlreadyDisabled
	}

	// Manipulate the gathered data
	out.PrintStandardSuccess()
	return

}
