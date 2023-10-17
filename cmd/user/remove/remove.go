package remove

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/user/complete"
	"github.com/itera-io/taikun-cli/utils/out"
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
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.UsersAPI.UsersDelete(context.TODO(), userID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("User", userID)

	return
}
