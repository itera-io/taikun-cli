package info

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/cmd/user/list"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var infoFields = list.ListFields

// NewCmdInfo creates and returns a cobra command for getting more Information on users.
// When called without an argument it gets the short User Info about the current user
// When called with one argument (user ID) it gets the long List information about the user with corresponding ID.
func NewCmdInfo() *cobra.Command {
	cmd := cobra.Command{
		Use:   "info [user-id]",
		Short: "Retrieve information about usertoken",
		Long:  "Retrieve information about usertoken (shows all the bound endpoints)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				infoFields.ShowAll()    // Long
				return listRun(args[0]) // List user by ID
			}
			// infoFields.ShowAll()
			return myInfoRun() // Info about current user
		},
	}

	complete.CompleteArgsWithUserID(&cmd)

	return &cmd
}

// myInfoRun calls the API and gets the info about the current user
func myInfoRun() (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.UsersAPI.UsersUserInfo(context.TODO()).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data.Data, infoFields)
}

// listRun calls the API and gets the info about user with userID
func listRun(userID string) (err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.UsersAPI.UsersList(context.TODO()).Id(userID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// User not found
	if len(data.Data) != 1 {
		return cmderr.ResourceNotFoundError("User", userID)
	}

	return out.PrintResult(data.Data[0], infoFields)
}
