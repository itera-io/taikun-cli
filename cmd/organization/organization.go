package organization

import (
	"taikun-cli/cmd/organization/list"

	"github.com/spf13/cobra"
)

func NewCmdOrganization() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "organization",
		Short:   "Manage organizations",
		Aliases: []string{"org"},
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
