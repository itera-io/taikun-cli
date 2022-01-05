package rule

import (
	"github.com/itera-io/taikun-cli/cmd/billing/rule/create"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/delete"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/list"
	"github.com/spf13/cobra"
)

func NewCmdRule() *cobra.Command {
	cmd := cobra.Command{
		Use:     "rule <command>",
		Short:   "Manage billing rules",
		Aliases: []string{"r"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
