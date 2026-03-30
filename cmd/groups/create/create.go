package create

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	AccountID int32
}

func NewCmdCreateGroup() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <GROUP_NAME>",
		Short: "Create a new group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createGroup(args[0], opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.AccountID, "account-id", "a", 0, "Account ID")
	_ = cmd.MarkFlagRequired("account-id")

	return &cmd
}

func createGroup(groupName string, opts CreateOptions) (err error) {
	// input parameters sanity check
	if opts.AccountID == 0 {
		return fmt.Errorf("account ID must be specified")
	}
	myApiClient := tk.NewClient()

	body := taikuncore.CreateGroupCommand{
		Name:      groupName,
		AccountId: opts.AccountID,
	}

	_, response, err := myApiClient.Client.GroupsAPI.GroupsCreate(context.TODO()).CreateGroupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
