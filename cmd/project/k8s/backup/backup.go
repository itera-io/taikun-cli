package backup

import (
	"github.com/itera-io/taikun-cli/cmd/project/k8s/backup/disable"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/backup/enable"
	"github.com/spf13/cobra"
)

func NewCmdBackup() *cobra.Command {
	cmd := cobra.Command{
		Use:   "backup <command>",
		Short: "Manage Kubernetes servers' backups",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enable.NewCmdEnable())

	return &cmd
}
