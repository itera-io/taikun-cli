package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "projectId",
		),
		field.NewVisible(
			"NAME", "projectName",
		),
	},
)

type ListOptions struct {
	UserID string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <user-id>",
		Short: "List a user's assigned projects",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return listRun(&opts)
		},
	}

	complete.CompleteArgsWithUserID(&cmd)

	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersListParams().WithV(api.Version)
	params = params.WithID(&opts.UserID)

	response, err := apiClient.Client.Users.UsersList(params, apiClient)
	if err != nil {
		return
	}

	if len(response.Payload.Data) != 1 {
		return cmderr.ResourceNotFoundError("User", opts.UserID)
	}

	return out.PrintResults(response.Payload.Data[0].BoundProjects, listFields)
}
