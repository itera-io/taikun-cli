package rule

import (
	"github.com/itera-io/taikun-cli/cmd/billing/rule/create"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/delete"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/label"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/list"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/organization"
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
	cmd.AddCommand(label.NewCmdLabel())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(organization.NewCmdOrganization())

	return &cmd
}
