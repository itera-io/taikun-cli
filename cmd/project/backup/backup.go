package backup

import (
	"github.com/itera-io/taikun-cli/cmd/project/backup/disable"
	"github.com/itera-io/taikun-cli/cmd/project/backup/enable"
	"github.com/itera-io/taikun-cli/cmd/project/backup/list"
	"github.com/itera-io/taikun-cli/cmd/project/backup/remove"
	"github.com/spf13/cobra"
)

func NewCmdBackup() *cobra.Command {
	cmd := cobra.Command{
		Use:   "backup <command>",
		Short: "Manage project's backup",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enable.NewCmdEnable())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return &cmd
}
