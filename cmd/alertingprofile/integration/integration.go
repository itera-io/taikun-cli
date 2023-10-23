package integration

import (
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration/add"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration/list"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration/remove"
	"github.com/spf13/cobra"
)

func NewCmdIntegration() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "integration <command>",
		Short:   "Manage an alerting profile's integrations",
		Aliases: []string{"int"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return cmd
}
