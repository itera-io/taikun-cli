package unbind

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/user_projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type UnbindOptions struct {
	UserID    string
	ProjectID int32
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := &cobra.Command{
		Use:   "unbind <user-id>",
		Short: "Unbind a user from a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return unbindRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(cmd, "project-id")

	return cmd
}

func unbindRun(opts *UnbindOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.BindProjectsCommand{
		UserID: opts.UserID,
		Projects: []*models.UpdateUserProjectDto{
			{
				ProjectID: opts.ProjectID,
				IsBound:   false,
			},
		},
	}

	params := user_projects.NewUserProjectsBindProjectsParams().WithV(api.Version).WithBody(body)
	_, err = apiClient.Client.UserProjects.UserProjectsBindProjects(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
