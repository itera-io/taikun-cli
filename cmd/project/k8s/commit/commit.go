package commit

import (
	"context"
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
		Short: "Commit changes to a project's Kubernetes servers",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return commitRun(&opts)
		},
	}

	return &cmd
}

func commitRun(opts *CommitOptions) (err error) {
	myApiClient := tk.NewClient()
	commitCommand := taikuncore.ProjectDeploymentCommitCommand{ProjectId: &opts.ProjectID}
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentCommit(context.TODO()).ProjectDeploymentCommitCommand(commitCommand).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
