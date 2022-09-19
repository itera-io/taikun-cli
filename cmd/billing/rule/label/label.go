package label

import (
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/add"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/list"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/remove"
	"github.com/spf13/cobra"
)

func NewCmdLabel() *cobra.Command {
	cmd := cobra.Command{
		Use:   "label",
		Short: "Manage a billing rule's labels",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return &cmd
}
