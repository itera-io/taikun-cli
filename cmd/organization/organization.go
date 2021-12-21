package organization

import (
	"github.com/itera-io/taikun-cli/cmd/organization/create"
	"github.com/itera-io/taikun-cli/cmd/organization/delete"
	"github.com/itera-io/taikun-cli/cmd/organization/list"

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
