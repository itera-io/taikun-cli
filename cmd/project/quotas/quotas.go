package quotas

import (
	"github.com/itera-io/taikun-cli/cmd/project/quotas/edit"
	"github.com/itera-io/taikun-cli/cmd/project/quotas/list"

	"github.com/spf13/cobra"
)

func NewCmdQuotas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quotas <command>",
		Short: "Manage project quotas",
	}

	cmd.AddCommand(edit.NewCmdEdit())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
