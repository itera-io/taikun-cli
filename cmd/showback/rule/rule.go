package rule

import (
	"github.com/itera-io/taikun-cli/cmd/showback/rule/create"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/delete"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/list"
	"github.com/spf13/cobra"
)

func NewCmdRule() *cobra.Command {
	cmd := cobra.Command{
		Use:     "rule <command>",
		Short:   "Manage showback rules",
		Aliases: []string{"r"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
