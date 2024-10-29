package label

import (
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/edit"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label/list"
	"github.com/spf13/cobra"
)

func NewCmdLabel() *cobra.Command {
	cmd := cobra.Command{
		Use:   "label",
		Short: "Manage a billing rule's labels",
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(edit.NewCmdEdit())

	return &cmd
}
