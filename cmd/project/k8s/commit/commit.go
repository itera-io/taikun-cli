package commit

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return commitRun(cmd, &opts)
		},
	}

	return &cmd
}

func commitRun(cmd *cobra.Command, opts *CommitOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	commitCommand := taikuncore.ProjectDeploymentCommitCommand{ProjectId: &opts.ProjectID}
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentCommit(ctx).ProjectDeploymentCommitCommand(commitCommand).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
