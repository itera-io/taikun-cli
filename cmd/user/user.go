package user

import (
	"taikun-cli/cmd/user/bind"
	"taikun-cli/cmd/user/list"
	"taikun-cli/cmd/user/unbind"

	"github.com/spf13/cobra"
)

func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user <command>",
		Short: "Manage users",
	}

	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(unbind.NewCmdUnbind())

	return cmd
}
