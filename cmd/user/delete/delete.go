package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <user-id>...",
		Short: "Delete one or more users",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdutils.DeleteMultipleStringID(args, deleteRun)
		},
	}

	return cmd
}

func deleteRun(id string) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDeleteParams().WithV(apiconfig.Version).WithID(id)
	_, _, err = apiClient.Client.Users.UsersDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("User", id)
	}

	return
}
