package delete

import (
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type DeleteOptions struct {
	UserIDs []string
}

func NewCmdDeleteUsers() *cobra.Command {
	opts := DeleteOptions{
		UserIDs: make([]string, 0),
	}

	cmd := cobra.Command{
		Use:   "delete <GROUP_ID>",
		Short: "Remove users from the group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return deleteUsersFromGroup(cmd, groupID, &opts)
		},
	}

	cmd.Flags().StringArrayVarP(&opts.UserIDs, "user-id", "u", nil, "User IDs")
	return &cmd
}

func deleteUsersFromGroup(cmd *cobra.Command, groupID int32, opts *DeleteOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// input parameters sanity check
	if len(opts.UserIDs) == 0 {
		return fmt.Errorf("no user IDs are specified")
	}
	myApiClient := tk.NewClient()

	body := *taikuncore.NewDeleteUserFromGroupCommand(groupID, opts.UserIDs)
	response, err := myApiClient.Client.GroupsAPI.GroupsDeleteUsers(ctx).DeleteUserFromGroupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
