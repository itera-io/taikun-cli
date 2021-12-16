package integration

import "github.com/spf13/cobra"

func NewCmdIntegration() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "integration <command>",
		Short: "Manage alerting integrations",
	}

	// TODO add subcommands

	return cmd
}
