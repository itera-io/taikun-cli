package organization

import "github.com/spf13/cobra"

func NewCmdOrganization() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "organization",
		Short:   "Manage organizations",
		Aliases: []string{"org"},
	}

	// TODO add subcommands

	return cmd
}
