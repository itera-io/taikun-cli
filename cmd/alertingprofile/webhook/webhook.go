package webhook

import (
	"taikun-cli/cmd/alertingprofile/webhook/add"
	"taikun-cli/cmd/alertingprofile/webhook/clear"
	"taikun-cli/cmd/alertingprofile/webhook/list"

	"github.com/spf13/cobra"
)

func NewCmdWebhook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook <command>",
		Short: "Manage alerting profile webhooks",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(clear.NewCmdClear())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
