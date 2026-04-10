package projects

import (
	"github.com/itera-io/taikun-cli/cmd/accounts/projects/info"
	"github.com/itera-io/taikun-cli/cmd/accounts/projects/list"
	"github.com/spf13/cobra"
)

func NewCmdProjects() *cobra.Command {
	cmd := cobra.Command{
		Use:   "projects <command>",
		Short: "Manage account projects",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(info.NewCmdInfo())

	return &cmd
}
