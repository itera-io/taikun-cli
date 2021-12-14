package flavor

import (
	"taikun-cli/cmd/flavor/all"
	"taikun-cli/cmd/flavor/bind"
	"taikun-cli/cmd/flavor/list"
	"taikun-cli/cmd/flavor/unbind"

	"github.com/spf13/cobra"
)

func NewCmdFlavor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flavor <command>",
		Short: "Get flavor info and manage project-flavor bindings",
	}

	cmd.AddCommand(all.NewCmdAll())
	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(unbind.NewCmdUnbind())

	return cmd
}
