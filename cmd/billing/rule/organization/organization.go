package organization

import (
	"github.com/itera-io/taikun-cli/cmd/billing/rule/organization/bind"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/organization/list"
	"github.com/itera-io/taikun-cli/cmd/billing/rule/organization/unbind"
	"github.com/spf13/cobra"
)

func NewCmdOrganization() *cobra.Command {
	cmd := cobra.Command{
		Use:     "organization <command>",
		Short:   "Manage a billing rule's organization bindings",
		Aliases: []string{"org"},
	}

	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(unbind.NewCmdUnbind())

	return &cmd
}
