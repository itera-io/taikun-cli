package delete

import (
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	UserID string
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := &cobra.Command{
		Use:   "delete <user-id>",
		Short: "Delete user",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return deleteRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	return cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDeleteParams().WithV(cmdutils.ApiVersion).WithID(opts.UserID)
	response, _, err := apiClient.Client.Users.UsersDelete(params, apiClient)
	if err == nil {
		cmdutils.PrettyPrint(response)
	}

	return
}
