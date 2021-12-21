package quotas

import (
	"github.com/itera-io/taikun-cli/cmd/project/quotas/list"
	"github.com/itera-io/taikun-cli/cmd/project/quotas/update"

	"github.com/spf13/cobra"
)

func NewCmdQuotas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotas <command>",
		Short: "Manage projects quotas",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(update.NewCmdUpdate())

	return cmd
}
