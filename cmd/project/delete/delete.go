package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	Force     bool
	ProjectID int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := &cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			opts.ProjectID = projectID
			return deleteRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Force, "force", "f", false, "Force delete")

	return cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteProjectCommand{
		IsForceDelete: opts.Force,
		ProjectID:     opts.ProjectID,
	}

	params := projects.NewProjectsDeleteParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	_, _, err = apiClient.Client.Projects.ProjectsDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Project", opts.ProjectID)
	}

	return
}
