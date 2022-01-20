package info

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/cmd/user/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

var infoFields = list.ListFields

func NewCmdInfo() *cobra.Command {
	cmd := cobra.Command{
		Use:   "info [user-id]",
		Short: "Retrieve a user's information",
		Long:  "Retrieve a user's information (yours if no user ID is specified)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				return infoRun(args[0])
			}
			infoFields.ShowAll()
			return myInfoRun()
		},
	}

	complete.CompleteArgsWithUserID(&cmd)

	return &cmd
}

func myInfoRun() (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDetailsParams().WithV(api.Version)
	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload.Data, infoFields)
	}

	return
}

func infoRun(userID string) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersListParams().WithV(api.Version)
	params = params.WithID(&userID)

	response, err := apiClient.Client.Users.UsersList(params, apiClient)
	if err != nil {
		return
	}
	if len(response.Payload.Data) != 1 {
		return cmderr.ResourceNotFoundError("User", userID)
	}

	out.PrintResult(response.Payload.Data[0], infoFields)

	return
}
