package quotas

import (
	"taikun-cli/cmd/project/quotas/list"
	"taikun-cli/cmd/project/quotas/update"

	"github.com/spf13/cobra"
)

func NewCmdQuotas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotas <command>",
		Short: "Manage project quotas",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(update.NewCmdUpdate())

	return cmd
}
