package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <user-token-name> ...",
		Short: "Delete one or more user tokens",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmdutils.DeleteMultipleStringID(args, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}
	complete.CompleteArgsWithUserTokenName(&cmd)
	return &cmd
}

func deleteRun(userTokenName string) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Get usertoken ID from usertoken name
	userTokenId, err := complete.UserTokenIDFromUserTokenName(userTokenName)
	if err != nil {
		return err
	}

	// Execute a query into the API + graceful exit
	_, err = myApiClient.Client.UserTokenAPI.UsertokenDelete(context.TODO(), userTokenId).Execute()
	if err == nil {
		out.PrintDeleteSuccess("User Token", userTokenName)
	} else {
		return err
	}

	return
}
