package rule

import (
	"github.com/itera-io/taikun-cli/cmd/showback/rule/add"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/list"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/remove"
	"github.com/spf13/cobra"
)

func NewCmdRule() *cobra.Command {
	cmd := cobra.Command{
		Use:     "rule <command>",
		Short:   "Manage showback rules",
		Aliases: []string{"r"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(label.NewCmdLabel())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return &cmd
}
