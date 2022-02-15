package user

import (
	"github.com/itera-io/taikun-cli/cmd/user/add"
	"github.com/itera-io/taikun-cli/cmd/user/delete"
	"github.com/itera-io/taikun-cli/cmd/user/info"
	"github.com/itera-io/taikun-cli/cmd/user/list"
	"github.com/itera-io/taikun-cli/cmd/user/project"
	"github.com/spf13/cobra"
)

func NewCmdUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user <command>",
		Short: "Manage users",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(project.NewCmdProject())

	return cmd
}
