package alertingprofile

import "github.com/spf13/cobra"

func NewCmdAlertingProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerting-profile",
		Short:   "Manage alerting profiles",
		Aliases: []string{"alert"},
	}

	// TODO add subcommands

	return cmd
}
