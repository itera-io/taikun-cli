package alertingprofile

import (
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/add"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/list"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/remove"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/unlock"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/webhook"
	"github.com/spf13/cobra"
)

func NewCmdAlertingProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerting-profile <command>",
		Short:   "Manage alerting profiles",
		Aliases: []string{"alp"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(integration.NewCmdIntegration())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())
	cmd.AddCommand(webhook.NewCmdWebhook())

	return cmd
}
