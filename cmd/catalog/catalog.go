package catalog

import (
	"github.com/itera-io/taikun-cli/cmd/catalog/app"
	"github.com/itera-io/taikun-cli/cmd/catalog/create"
	"github.com/itera-io/taikun-cli/cmd/catalog/delete"
	"github.com/itera-io/taikun-cli/cmd/catalog/list"
	"github.com/itera-io/taikun-cli/cmd/catalog/lock"
	"github.com/itera-io/taikun-cli/cmd/catalog/makedefault"
	"github.com/itera-io/taikun-cli/cmd/catalog/project"
	"github.com/itera-io/taikun-cli/cmd/catalog/unlock"
	"github.com/spf13/cobra"
)

func NewCmdCatalog() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "catalog  <command>",
		Short:   "Manage catalogs in Taikun",
		Aliases: []string{"catalogs", "cat", "cats"},
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(project.NewCmdCatalogProject())
	cmd.AddCommand(makedefault.NewCmdMakedefault())
	cmd.AddCommand(app.NewCmdApp())
	cmd.AddCommand(create.NewCmdCreatecatalog())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
