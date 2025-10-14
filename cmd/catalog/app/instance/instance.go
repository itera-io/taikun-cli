package instance

import (
	"github.com/itera-io/taikun-cli/cmd/catalog/app/instance/cancel"
	"github.com/itera-io/taikun-cli/cmd/catalog/app/instance/install"
	"github.com/itera-io/taikun-cli/cmd/catalog/app/instance/list"
	syncpackage "github.com/itera-io/taikun-cli/cmd/catalog/app/instance/sync"
	"github.com/itera-io/taikun-cli/cmd/catalog/app/instance/uninstall"
	"github.com/spf13/cobra"
)

func NewCmdInstance() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "instance <command>",
		Short:   "Install and uninstall catalog applications",
		Aliases: []string{"instances", "inst"},
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(syncpackage.NewCmdSync())
	cmd.AddCommand(uninstall.NewCmdUninstall())
	cmd.AddCommand(install.NewCmdInstall())
	cmd.AddCommand(cancel.NewCmdCancel())

	return cmd
}
