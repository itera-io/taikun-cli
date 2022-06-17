package backupsource

import (
	"github.com/itera-io/taikun-cli/cmd/project/backupsource/add"
	"github.com/spf13/cobra"
)

func NewCmdBackupSource() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backup-source <command>",
		Short: "Manage a project's backup sources",
	}

	cmd.AddCommand(add.NewCmdAdd())
	return cmd
}
