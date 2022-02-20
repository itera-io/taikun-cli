package unlock

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <project-id>",
		Short: "Unlock a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unlockRun(projectID)
		},
	}

	return cmd
}

func unlockRun(projectID int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := projects.NewProjectsLockManagerParams().WithV(api.Version)
	params = params.WithMode(&types.UnlockedMode).WithID(&projectID)

	_, err = apiClient.Client.Projects.ProjectsLockManager(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
