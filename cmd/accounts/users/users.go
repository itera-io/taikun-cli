package users

import (
	"github.com/itera-io/taikun-cli/cmd/accounts/users/info"
	"github.com/itera-io/taikun-cli/cmd/accounts/users/list"
	"github.com/spf13/cobra"
)

func NewCmdUsers() *cobra.Command {
	cmd := cobra.Command{
		Use:   "users <command>",
		Short: "Manage account users",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(info.NewCmdInfo())

	return &cmd
}
