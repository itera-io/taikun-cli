package delete

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <user-id>",
		Short: "Delete user",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteRun(args[0])
		},
		Args: cobra.ExactArgs(1),
	}

	return cmd
}

func deleteRun(id string) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDeleteParams().WithV(cmdutils.ApiVersion).WithID(id)
	_, _, err = apiClient.Client.Users.UsersDelete(params, apiClient)
	if err == nil {
		fmt.Println("User deleted")
	}

	return
}
