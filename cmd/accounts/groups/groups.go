package groups

import (
	"github.com/itera-io/taikun-cli/cmd/accounts/groups/info"
	"github.com/itera-io/taikun-cli/cmd/accounts/groups/list"
	"github.com/spf13/cobra"
)

func NewCmdGroups() *cobra.Command {
	cmd := cobra.Command{
		Use:   "groups <command>",
		Short: "Manage account groups",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(info.NewCmdInfo())

	return &cmd
}
