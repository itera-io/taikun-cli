package worker

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type WorkerOptions struct {
	ProjectID         int32
	EnableWorkerSpot  bool
	DisableWorkerSpot bool
}

func NewCmdWorker() *cobra.Command {
	var opts WorkerOptions

	cmd := cobra.Command{
		Use:   "worker <project-id>",
		Short: "Disable or enable the project's spot Kubernetes worker support",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return workerRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.EnableWorkerSpot, "enable", "e", false, "Enable spot Kubernetes worker support (spot workers)")
	cmd.Flags().BoolVarP(&opts.DisableWorkerSpot, "disable", "d", false, "Disable spot Kubernetes worker support (spot workers)")
	cmd.MarkFlagsOneRequired("enable", "disable")
	cmd.MarkFlagsMutuallyExclusive("enable", "disable")

	return &cmd
}

func workerRun(opts *WorkerOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.SpotWorkerOperationCommand{
		Id: &opts.ProjectID,
	}

	if opts.EnableWorkerSpot {
		body.SetMode("enable")
	} else if opts.DisableWorkerSpot {
		body.SetMode("disable")
	} else {
		return fmt.Errorf("unknown mode. Either disable or enable")
	}

	_, response, err := myApiClient.Client.ProjectsAPI.ProjectsToggleSpotWorkers(context.TODO()).SpotWorkerOperationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
