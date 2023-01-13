package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/usertoken/complete"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/user_token"
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
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	userTokenId, err := complete.UserTokenIDFromUserTokenName(userTokenName)
	if err != nil {
		return
	}

	params := user_token.NewUserTokenDeleteParams().WithV(taikungoclient.Version).WithID(userTokenId)

	_, err = apiClient.Client.UserToken.UserTokenDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("User Token", userTokenName)
	}

	return
}
