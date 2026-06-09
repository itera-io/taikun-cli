package full

import (
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return fullRun(cmd, &opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.EnableFullSpot, "enable", "e", false, "Enable full spot Kubernetes support (spot controlplane and workers)")
	cmd.Flags().BoolVarP(&opts.DisableFullSpot, "disable", "d", false, "Disable full spot Kubernetes support (spot controlplane and workers)")
	cmd.MarkFlagsOneRequired("enable", "disable")
	cmd.MarkFlagsMutuallyExclusive("enable", "disable")

	return &cmd
}

func fullRun(cmd *cobra.Command, opts *FullOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	body := taikuncore.FullSpotOperationCommand{
		Id: &opts.ProjectID,
	}

	if opts.EnableFullSpot {
		body.SetMode("enable")
	} else if opts.DisableFullSpot {
		body.SetMode("disable")
	} else {
		return fmt.Errorf("unknown mode. Either disable or enable")
	}

	response, err := myApiClient.Client.ProjectsAPI.ProjectsToggleFullSpot(ctx).FullSpotOperationCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
