package repair

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type RepairOptions struct {
	ProjectID int32
}

func NewCmdRepair() *cobra.Command {
	var opts RepairOptions

	cmd := cobra.Command{
		Use:   "repair <project-id>",
		Short: "Repair a project's Kubernetes servers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return repairRun(cmd, &opts)
		},
	}

	return &cmd
}

func repairRun(cmd *cobra.Command, opts *RepairOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	var body taikuncore.ProjectDeploymentRepairCommand
	body.SetProjectId(opts.ProjectID)
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentRepair(ctx).ProjectDeploymentRepairCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
