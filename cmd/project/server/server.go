package server

import (
	"github.com/itera-io/taikun-cli/cmd/project/server/add"
	"github.com/itera-io/taikun-cli/cmd/project/server/delete"
	"github.com/itera-io/taikun-cli/cmd/project/server/list"
	"github.com/itera-io/taikun-cli/cmd/project/server/reboot"
	"github.com/itera-io/taikun-cli/cmd/project/server/status"
	"github.com/spf13/cobra"
)

func NewCmdServer() *cobra.Command {
	cmd := cobra.Command{
		Use:   "server <command>",
		Short: "Manage servers",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(reboot.NewCmdReboot())
	cmd.AddCommand(status.NewCmdStatus())

	return &cmd
}
