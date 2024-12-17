package app

import (
	"github.com/itera-io/taikun-cli/cmd/catalog/app/bind"
	"github.com/itera-io/taikun-cli/cmd/catalog/app/instance"
	"github.com/itera-io/taikun-cli/cmd/catalog/app/list"
	"github.com/itera-io/taikun-cli/cmd/catalog/app/unbind"
	"github.com/spf13/cobra"
)

func NewCmdApp() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "application <command>",
		Short:   "Manage and catalog applications",
		Aliases: []string{"applications", "app", "apps"},
	}

	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(unbind.NewCmdUnbind())
	cmd.AddCommand(instance.NewCmdInstance())

	return cmd
}
