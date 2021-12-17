package alertingprofile

import (
	"taikun-cli/cmd/alertingprofile/create"
	"taikun-cli/cmd/alertingprofile/delete"
	"taikun-cli/cmd/alertingprofile/integration"
	"taikun-cli/cmd/alertingprofile/list"
	"taikun-cli/cmd/alertingprofile/lock"
	"taikun-cli/cmd/alertingprofile/unlock"
	"taikun-cli/cmd/alertingprofile/webhook"

	"github.com/spf13/cobra"
)

func NewCmdAlertingProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerting-profile <command>",
		Short:   "Manage alerting profiles",
		Aliases: []string{"alert"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(integration.NewCmdIntegration())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(webhook.NewCmdWebhook())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
