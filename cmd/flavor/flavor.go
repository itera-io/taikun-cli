package flavor

import "github.com/spf13/cobra"

func NewCmdFlavor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "flavor <command>",
		Short: "Get flavor info and manage project-flavor bindings",
	}

	// TODO add subcommands

	return cmd
}
