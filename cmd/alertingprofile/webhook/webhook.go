package webhook

import "github.com/spf13/cobra"

func NewCmdWebhook() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "webhook <command>",
		Short: "Manage alerting profile webhooks",
	}

	// TODO add subcommands

	return cmd
}
