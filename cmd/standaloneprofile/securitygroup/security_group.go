package securitygroup

import (
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/securitygroup/add"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/securitygroup/list"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/securitygroup/remove"
	"github.com/spf13/cobra"
)

func NewCmdSecurityGroup() *cobra.Command {
	cmd := cobra.Command{
		Use:     "security-group <command>",
		Short:   "Manage securitygroups",
		Aliases: []string{"group"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return &cmd
}
