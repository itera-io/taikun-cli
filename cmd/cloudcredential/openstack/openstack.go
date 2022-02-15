package openstack

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/add"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/list"
	"github.com/spf13/cobra"
)

func NewCmdOpenstack() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "openstack <command>",
		Short:   "Manage OpenStack cloud credentials",
		Aliases: []string{"os"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
