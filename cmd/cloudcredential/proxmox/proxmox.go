package proxmox

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/proxmox/add"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/proxmox/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/proxmox/list"
	"github.com/spf13/cobra"
)

func NewCmdProxmox() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxmox <command>",
		Short: "Manage Proxmox cloud credentials",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
