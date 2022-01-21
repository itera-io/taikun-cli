package vm

import "github.com/spf13/cobra"

func NewCmdVm() *cobra.Command {
	cmd := cobra.Command{
		Use:   "vm <command>",
		Short: "Manage a project's standalone VMs",
	}

	// FIXME

	return &cmd
}
