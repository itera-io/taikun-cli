package full

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type FullOptions struct {
	ProjectID       int32
	EnableFullSpot  bool
	DisableFullSpot bool
}

func NewCmdFull() *cobra.Command {
	var opts FullOptions

	cmd := cobra.Command{
		Use:   "full <project-id>",
		Short: "Disable or enable the project's full spot Kubernetes support",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return fullRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.EnableFullSpot, "enable", "e", false, "Enable full spot Kubernetes support (spot controlplane and workers)")
	cmd.Flags().BoolVarP(&opts.DisableFullSpot, "disable", "d", false, "Disable full spot Kubernetes support (spot controlplane and workers)")
	cmd.MarkFlagsOneRequired("enable", "disable")
	cmd.MarkFlagsMutuallyExclusive("enable", "disable")

	return &cmd
}

func fullRun(opts *FullOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.FullSpotOperationCommand{
		Id: &opts.ProjectID,
	}

	if opts.EnableFullSpot {
		body.SetMode("enable")
	} else if opts.DisableFullSpot {
		body.SetMode("disable")
	} else {
		return fmt.Errorf("Unknown mode. Either disable or enable.")
	}

	_, response, err := myApiClient.Client.ProjectsAPI.ProjectsToggleFullSpot(context.TODO()).FullSpotOperationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
