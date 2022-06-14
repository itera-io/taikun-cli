package restore

import (
	"github.com/itera-io/taikun-cli/cmd/project/restore/add"
	"github.com/spf13/cobra"
)

func NewCmdRestore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore <command>",
		Short: "Manage a project's backup restores",
	}

	cmd.AddCommand(add.NewCmdAdd())

	return cmd
}
