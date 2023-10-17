package remove

import (
	"context"
	"errors"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"os"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	Force     bool
	ProjectID int32
}

func NewCmdDelete() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "delete <project-id>...",
		Short: "Delete one or more projects",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			optsList := make([]*DeleteOptions, len(args))
			for i, arg := range args {
				projectID, err := types.Atoi32(arg)
				if err != nil {
					return cmderr.ErrIDArgumentNotANumber
				}
				optsList[i] = &DeleteOptions{
					Force:     force,
					ProjectID: projectID,
				}
			}
			return deleteMultiple(optsList)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force delete")

	return cmd
}

func deleteMultiple(optsList []*DeleteOptions) error {
	errorOccured := false

	for _, opts := range optsList {
		if err := deleteRun(opts); err != nil {
			errorOccured = true

			fmt.Fprintln(os.Stderr, err)
		}
	}

	if errorOccured {
		fmt.Fprintln(os.Stderr)
		return errors.New("Failed to delete one or more projects")
	}

	return nil
}

func deleteRun(opts *DeleteOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.DeleteProjectCommand{
		ProjectId:     &opts.ProjectID,
		IsForceDelete: &opts.Force,
	}
	request, err := myApiClient.Client.ProjectsAPI.ProjectsDelete(context.TODO()).DeleteProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(request, err)
	}
	out.PrintDeleteSuccess("Project", opts.ProjectID)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.DeleteProjectCommand{
			IsForceDelete: opts.Force,
			ProjectID:     opts.ProjectID,
		}

		params := projects.NewProjectsDeleteParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, _, err = apiClient.Client.Projects.ProjectsDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Project", opts.ProjectID)
		}
	*/

	return
}
