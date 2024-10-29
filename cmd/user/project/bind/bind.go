package bind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	UserID    string
	ProjectID int32
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := cobra.Command{
		Use:   "bind <user-id>",
		Short: "Bind a user to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return bindRun(&opts)
		},
	}

	complete.CompleteArgsWithUserID(&cmd)

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "project-id")

	return &cmd
}

// bindRun calls the API at /usersprojects/bindprojects and binds a user to a project.
// Both identified by ID. If user and project are already bound if fails.
func bindRun(opts *BindOptions) (err error) {
	myApiClient := tk.NewClient()

	// Create the body for the request
	body := []int32{opts.ProjectID}

	// Send the request and process response
	//response, err := myApiClient.Client.UserProjectsAPI.UserprojectsBindProjects(context.TODO()).BindProjectsCommand(body).Execute()
	response, err := myApiClient.Client.UsersAPI.UsersAddUserProjects(context.TODO(), opts.UserID).RequestBody(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return
}
