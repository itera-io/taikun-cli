package unbind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

type UnbindOptions struct {
	UserID    string
	ProjectID int32
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := cobra.Command{
		Use:   "unbind <user-id>",
		Short: "Unbind a user from a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return unbindRun(&opts)
		},
	}

	complete.CompleteArgsWithUserID(&cmd)

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "project-id")

	return &cmd
}

// unbindRun calls the API at /usersprojects/bindprojects and unbinds a user from a project.
// Both identified by ID. If user and project are already unbound if fails.
func unbindRun(opts *UnbindOptions) (err error) {
	myApiClient := tk.NewClient()

	// Create the body for the request
	body := []int32{opts.ProjectID}

	// Send the request and process response
	response, err := myApiClient.Client.UsersAPI.UsersDeleteUserProjects(context.TODO(), opts.UserID).RequestBody(body).Execute()

	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return
}
