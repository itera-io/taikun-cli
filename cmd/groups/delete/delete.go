package delete

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdDeleteGroup() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <GROUP_ID>",
		Short: "Delete group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return deleteGroup(cmd, groupID)
		},
	}

	return &cmd
}

func deleteGroup(cmd *cobra.Command, groupID int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.GroupsAPI.GroupsDelete(ctx, groupID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("group", groupID)
	return nil
}
