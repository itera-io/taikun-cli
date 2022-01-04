package label

import (
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/add"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/clear"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/spf13/cobra"
)

func NewCmdLabel() *cobra.Command {
	cmd := cobra.Command{
		Use:   "label <command>",
		Short: "Manage showback rule labels",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(clear.NewCmdClear())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
