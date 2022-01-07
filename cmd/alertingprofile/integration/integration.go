package integration

import (
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration/create"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration/delete"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile/integration/list"

	"github.com/spf13/cobra"
)

func NewCmdIntegration() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "integration <command>",
		Short:   "Manage alerting integrations",
		Aliases: []string{"int"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
