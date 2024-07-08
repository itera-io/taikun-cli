package repair

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
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
			return repairRun(&opts)
		},
	}

	return &cmd
}

func repairRun(opts *RepairOptions) (err error) {
	myApiClient := tk.NewClient()
	var body taikuncore.ProjectDeploymentRepairCommand
	body.SetProjectId(opts.ProjectID)
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentRepair(context.TODO()).ProjectDeploymentRepairCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
