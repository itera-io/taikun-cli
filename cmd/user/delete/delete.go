package delete

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <user-id>...",
		Short: "Delete one or more users",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdutils.DeleteMultipleStringID(args, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	complete.CompleteArgsWithUserID(&cmd)

	return &cmd
}

func deleteRun(userID string) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDeleteParams().WithV(api.Version).WithID(userID)
	_, _, err = apiClient.Client.Users.UsersDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("User", userID)
	}

	return
}
