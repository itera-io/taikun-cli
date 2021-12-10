package bind

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/user_projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	Username  string
	ProjectID int
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := &cobra.Command{
		Use:   "bind",
		Short: "Bind a user to a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return bindRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Username, "username", "u", "", "Username (required)")
	cmdutils.MarkFlagRequired(cmd, "username")

	cmd.Flags().IntVarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(cmd, "project-id")

	return cmd
}

func bindRun(opts *BindOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.BindProjectsCommand{
		UserName: opts.Username,
		Projects: []*models.UpdateUserProjectDto{
			{
				ProjectID: int32(opts.ProjectID),
				IsBound:   true,
			},
		},
	}

	params := user_projects.NewUserProjectsBindProjectsParams().WithV(cmdutils.ApiVersion).WithBody(body)
	response, err := apiClient.Client.UserProjects.UserProjectsBindProjects(params, apiClient)
	if err == nil {
		fmt.Println(response.Payload)
	}

	return
}
