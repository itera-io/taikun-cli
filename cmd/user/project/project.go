package project

import (
	"github.com/itera-io/taikun-cli/cmd/user/project/bind"
	"github.com/itera-io/taikun-cli/cmd/user/project/list"
	"github.com/itera-io/taikun-cli/cmd/user/project/unbind"
	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := cobra.Command{
		Use:   "project <command>",
		Short: "Manage a user's assigned projects",
	}

	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(unbind.NewCmdUnbind())

	return &cmd
}
