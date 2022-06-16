package restore

import (
	"github.com/itera-io/taikun-cli/cmd/project/restore/add"
	"github.com/itera-io/taikun-cli/cmd/project/restore/list"
	"github.com/spf13/cobra"
)

func NewCmdRestore() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore <command>",
		Short: "Manage a project's backup restores",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	return cmd
}
