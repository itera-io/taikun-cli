package webhook

import (
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/webhook/add"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/webhook/clear"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/webhook/list"

	"github.com/spf13/cobra"
)

func NewCmdWebhook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook <command>",
		Short: "Manage an alerting profile's webhooks",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(clear.NewCmdClear())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
