package robot

import (
	"github.com/itera-io/taikun-cli/cmd/robot/checker"
	"github.com/itera-io/taikun-cli/cmd/robot/create"
	"github.com/itera-io/taikun-cli/cmd/robot/delete"
	"github.com/itera-io/taikun-cli/cmd/robot/details"
	"github.com/itera-io/taikun-cli/cmd/robot/list"
	"github.com/itera-io/taikun-cli/cmd/robot/regenerate"
	scope_list "github.com/itera-io/taikun-cli/cmd/robot/scope-list"
	"github.com/itera-io/taikun-cli/cmd/robot/status"
	"github.com/itera-io/taikun-cli/cmd/robot/update"
	updatescope "github.com/itera-io/taikun-cli/cmd/robot/update-scope"
	"github.com/spf13/cobra"
)

func NewCmdRobot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "robot  <command>",
		Short:   "Manage robot users in Taikun",
		Aliases: []string{"robot", "rbt", "r"},
	}

	cmd.AddCommand(create.NewCmdCreateRobot())
	cmd.AddCommand(list.NewCmdListRobots())
	cmd.AddCommand(delete.NewCmdDeleteRobot())
	cmd.AddCommand(scope_list.NewCmdScopeList())
	cmd.AddCommand(regenerate.NewCmdRobotRegenerateTokens())
	cmd.AddCommand(checker.NewCmdCreateRobotChecker())
	cmd.AddCommand(status.NewCmdRobotStatus())
	cmd.AddCommand(update.NewCmdUpdate())
	cmd.AddCommand(updatescope.NewCmdUpdateScope())
	cmd.AddCommand(details.NewCmdDetails())
	//cmd.AddCommand(delete.NewCmdDeleteGroup())
	//cmd.AddCommand(update.NewCmdUpdateGroup())
	//cmd.AddCommand(check.NewCmdCheckDuplicateEntity())
	//cmd.AddCommand(users.NewCmdGroupsUsers())
	//cmd.AddCommand(organizations.NewCmdGroupsOrganizations())
	return cmd
}
