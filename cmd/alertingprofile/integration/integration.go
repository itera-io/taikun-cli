package integration

import (
	"taikun-cli/cmd/alertingprofile/integration/delete"
	"taikun-cli/cmd/alertingprofile/integration/list"

	"github.com/spf13/cobra"
)

func NewCmdIntegration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "integration <command>",
		Short: "Manage alerting integrations",
	}

	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
