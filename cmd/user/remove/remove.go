package remove

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <user-id>...",
		Short: "Delete one or more users",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := cmdutils.APIContext(cmd)
			defer cancel()
			return cmdutils.DeleteMultipleStringID(args, func(id string) error {
				return deleteRun(ctx, id)
			})
		},
		Aliases: cmdutils.DeleteAliases,
	}

	complete.CompleteArgsWithUserID(&cmd)

	return &cmd
}

func deleteRun(ctx context.Context, userID string) (err error) {
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.UsersAPI.UsersDelete(ctx, userID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("User", userID)

	return
}
