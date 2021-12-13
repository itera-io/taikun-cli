package flavor

import (
	"taikun-cli/cmd/flavor/bind"

	"github.com/spf13/cobra"
)

func NewCmdFlavor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flavor <command>",
		Short: "Get flavor info and manage project-flavor bindings",
	}

	cmd.AddCommand(bind.NewCmdBind())

	return cmd
}
