package add

import (
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type AddOptions struct {
	UserIDs []string
}

func NewCmdAddUser() *cobra.Command {
	opts := AddOptions{
		UserIDs: make([]string, 0),
	}

	cmd := cobra.Command{
		Use:   "add <GROUP_ID>",
		Short: "Add users to the group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return addUsersToGroup(cmd, groupID, &opts)
		},
	}

	cmd.Flags().StringArrayVarP(&opts.UserIDs, "user-id", "u", nil, "User IDs")
	return &cmd
}

func addUsersToGroup(cmd *cobra.Command, groupID int32, opts *AddOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// input parameters sanity check
	if len(opts.UserIDs) == 0 {
		return fmt.Errorf("no user IDs are specified")
	}
	myApiClient := tk.NewClient()

	body := make([]taikuncore.CreateGroupUserDto, 0)
	for _, userID := range opts.UserIDs {
		body = append(body, *taikuncore.NewCreateGroupUserDto(*taikuncore.NewNullableString(&userID)))
	}

	response, err := myApiClient.Client.GroupsAPI.GroupsAddUsers(ctx, groupID).CreateGroupUserDto(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
