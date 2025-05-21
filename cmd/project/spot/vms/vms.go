package vms

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type VmsOptions struct {
	ProjectID      int32
	EnableVmsSpot  bool
	DisableVmsSpot bool
}

func NewCmdVms() *cobra.Command {
	var opts VmsOptions

	cmd := cobra.Command{
		Use:   "vms <project-id>",
		Short: "Disable or enable the project's spot for standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return vmsRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.EnableVmsSpot, "enable", "e", false, "Enable spot for standalone VMs")
	cmd.Flags().BoolVarP(&opts.DisableVmsSpot, "disable", "d", false, "Disable spot for standalone VMs")
	cmd.MarkFlagsOneRequired("enable", "disable")
	cmd.MarkFlagsMutuallyExclusive("enable", "disable")

	return &cmd
}

func vmsRun(opts *VmsOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.SpotVmOperationCommand{
		Id: &opts.ProjectID,
	}

	if opts.EnableVmsSpot {
		body.SetMode("enable")
	} else if opts.DisableVmsSpot {
		body.SetMode("disable")
	} else {
		return fmt.Errorf("unknown mode. Either disable or enable")
	}

	response, err := myApiClient.Client.ProjectsAPI.ProjectsToggleSpotVms(context.TODO()).SpotVmOperationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
