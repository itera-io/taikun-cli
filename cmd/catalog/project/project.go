package project

import (
	"github.com/itera-io/taikun-cli/cmd/catalog/project/bind"
	"github.com/itera-io/taikun-cli/cmd/catalog/project/list"
	"github.com/itera-io/taikun-cli/cmd/catalog/project/unbind"
	"github.com/spf13/cobra"
)

func NewCmdCatalogProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project <command>",
		Short:   "Manage catalogs projects in Taikun",
		Aliases: []string{"projects", "proj"},
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(unbind.NewCmdUnbind())
	return cmd
}
