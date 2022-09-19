package organization

import (
	"github.com/itera-io/taikun-cli/cmd/organization/add"
	"github.com/itera-io/taikun-cli/cmd/organization/info"
	"github.com/itera-io/taikun-cli/cmd/organization/list"
	"github.com/itera-io/taikun-cli/cmd/organization/remove"
	"github.com/spf13/cobra"
)

func NewCmdOrganization() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "organization",
		Short:   "Manage organizations",
		Aliases: []string{"org"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return cmd
}
