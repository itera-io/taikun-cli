package organization

import (
	"taikun-cli/cmd/organization/create"
	"taikun-cli/cmd/organization/delete"
	"taikun-cli/cmd/organization/list"

	"github.com/spf13/cobra"
)

func NewCmdOrganization() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "organization",
		Short:   "Manage organizations",
		Aliases: []string{"org"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
