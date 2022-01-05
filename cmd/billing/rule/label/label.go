package label

import (
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/add"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/delete"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/list"
	"github.com/spf13/cobra"
)

func NewCmdLabel() *cobra.Command {
	cmd := cobra.Command{
		Use:   "label",
		Short: "Manage labels",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
