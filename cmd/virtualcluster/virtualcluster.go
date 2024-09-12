package virtualcluster

import (
	"github.com/itera-io/taikun-cli/cmd/virtualcluster/list"
	"github.com/spf13/cobra"
)

func NewCmdVirtualcluster() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "virtualcluster <command>",
		Short:   "Manage virtual cluster projects",
		Aliases: []string{"vc"},
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
