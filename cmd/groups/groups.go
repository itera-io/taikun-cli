package groups

import (
	"github.com/itera-io/taikun-cli/cmd/groups/check"
	"github.com/itera-io/taikun-cli/cmd/groups/create"
	"github.com/itera-io/taikun-cli/cmd/groups/delete"
	"github.com/itera-io/taikun-cli/cmd/groups/list"
	"github.com/itera-io/taikun-cli/cmd/groups/organizations"
	"github.com/itera-io/taikun-cli/cmd/groups/update"
	"github.com/itera-io/taikun-cli/cmd/groups/users"
	"github.com/spf13/cobra"
)

func NewCmdGroups() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "groups  <command>",
		Short:   "Manage groups in Taikun",
		Aliases: []string{"groups", "grp", "G"},
	}

	cmd.AddCommand(create.NewCmdCreateGroup())
	cmd.AddCommand(list.NewCmdListGroups())
	cmd.AddCommand(delete.NewCmdDeleteGroup())
	cmd.AddCommand(update.NewCmdUpdateGroup())
	cmd.AddCommand(check.NewCmdCheckDuplicateEntity())
	cmd.AddCommand(users.NewCmdGroupsUsers())
	cmd.AddCommand(organizations.NewCmdGroupsOrganizations())
	return cmd
}
