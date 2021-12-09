package bind

import (
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/user_projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	UserID    string
	ProjectID int
}

func NewCmdBind(apiClient *api.Client) *cobra.Command {
	var opts BindOptions

	cmd := &cobra.Command{
		Use:   "bind",
		Short: "Bind a user to a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return bindRun(&opts, apiClient)
		},
	}

	cmd.Flags().StringVarP(&opts.UserID, "user-id", "u", "", "User ID (required)")
	cmd.MarkFlagRequired("user-id")

	cmd.Flags().IntVarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmd.MarkFlagRequired("project-id")

	return cmd
}

func bindRun(opts *BindOptions, apiClient *api.Client) (err error) {
	body := &models.BindProjectsCommand{
		UserID: opts.UserID,
		Projects: []*models.UpdateUserProjectDto{
			{
				ProjectID: int32(opts.ProjectID),
				IsBound:   true,
			},
		},
	}

	params := user_projects.NewUserProjectsBindProjectsParams().WithV(cmdutils.ApiVersion).WithBody(body)
	_, err = apiClient.Client.UserProjects.UserProjectsBindProjects(params, apiClient)
	return
}
