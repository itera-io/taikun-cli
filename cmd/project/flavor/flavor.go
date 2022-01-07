package flavor

import (
	"github.com/itera-io/taikun-cli/cmd/project/flavor/bind"
	"github.com/itera-io/taikun-cli/cmd/project/flavor/list"
	"github.com/itera-io/taikun-cli/cmd/project/flavor/unbind"

	"github.com/spf13/cobra"
)

func NewCmdFlavor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flavor <command>",
		Short: "Manage a project's flavors",
	}

	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(unbind.NewCmdUnbind())

	return cmd
}
