package vsphere

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/vsphere/add"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/vsphere/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/vsphere/list"
	"github.com/spf13/cobra"
)

func NewCmdVsphere() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vsphere <command>",
		Short: "Manage vSphere cloud credentials",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
