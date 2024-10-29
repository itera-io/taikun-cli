package commit

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CommitOptions struct {
	ProjectID int32
}

func NewCmdCommit() *cobra.Command {
	var opts CommitOptions

	cmd := cobra.Command{
		Use:   "commit <project-id>",
		Short: "Commit changes to a project's standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return commitRun(&opts)
		},
	}

	return &cmd
}

func commitRun(opts *CommitOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.DeploymentCommitVmCommand{
		ProjectId: &opts.ProjectID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentCommitVm(context.TODO()).DeploymentCommitVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
