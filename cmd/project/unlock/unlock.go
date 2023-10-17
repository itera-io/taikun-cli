package unlock

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	body := taikuncore.ProjectLockManagerCommand{
		Id:   &projectID,
		Mode: *taikuncore.NewNullableString(&types.UnlockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.ProjectsAPI.ProjectsLockManager(context.TODO()).ProjectLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	out.PrintStandardSuccess()
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := projects.NewProjectsLockManagerParams().WithV(taikungoclient.Version)
		params = params.WithMode(&types.UnlockedMode).WithID(&projectID)

		_, err = apiClient.Client.Projects.ProjectsLockManager(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
