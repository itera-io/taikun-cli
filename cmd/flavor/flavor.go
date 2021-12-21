package flavor

import (
	"github.com/itera-io/taikun-cli/cmd/flavor/all"
	"github.com/itera-io/taikun-cli/cmd/flavor/bind"
	"github.com/itera-io/taikun-cli/cmd/flavor/list"
	"github.com/itera-io/taikun-cli/cmd/flavor/unbind"

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
