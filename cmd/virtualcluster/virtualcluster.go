package virtualcluster

import (
	"github.com/itera-io/taikun-cli/cmd/virtualcluster/add"
	"github.com/itera-io/taikun-cli/cmd/virtualcluster/list"
	"github.com/itera-io/taikun-cli/cmd/virtualcluster/remove"
	"github.com/spf13/cobra"
)

func NewCmdVirtualcluster() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "virtualcluster <command>",
		Short:   "Manage virtual cluster projects",
		Aliases: []string{"vc"},
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(add.NewCmdAdd())

	return cmd
}
