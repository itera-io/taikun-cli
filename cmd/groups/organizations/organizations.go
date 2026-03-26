package organizations

import (
	"github.com/itera-io/taikun-cli/cmd/groups/organizations/add"
	"github.com/itera-io/taikun-cli/cmd/groups/organizations/delete"
	"github.com/itera-io/taikun-cli/cmd/groups/organizations/update"
	"github.com/spf13/cobra"
)

func NewCmdGroupsOrganizations() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "organizations <command>",
		Short:   "Manage organizations within groups in Taikun",
		Aliases: []string{"organizations", "orgs", "Os"},
	}

	cmd.AddCommand(add.NewCmdAddOrganizations())
	cmd.AddCommand(update.NewCmdUpdateOrganization())
	cmd.AddCommand(delete.NewCmdDeleteOrganizations())
	return cmd
}
