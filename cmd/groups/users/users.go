package users

import (
	"github.com/itera-io/taikun-cli/cmd/groups/users/add"
	"github.com/itera-io/taikun-cli/cmd/groups/users/delete"
	"github.com/spf13/cobra"
)

func NewCmdGroupsUsers() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "users <command>",
		Short:   "Manage users within groups in Taikun",
		Aliases: []string{"users", "u"},
	}

	cmd.AddCommand(add.NewCmdAddUser())
	cmd.AddCommand(delete.NewCmdDeleteUsers())
	return cmd
}
