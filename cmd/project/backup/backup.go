package backup

import (
	"github.com/itera-io/taikun-cli/cmd/project/backup/disable"
	"github.com/itera-io/taikun-cli/cmd/project/backup/enable"
	"github.com/spf13/cobra"
)

func NewCmdBackup() *cobra.Command {
	cmd := cobra.Command{
		Use:   "backup <command>",
		Short: "Manage project's backup",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enable.NewCmdEnable())

	return &cmd
}
