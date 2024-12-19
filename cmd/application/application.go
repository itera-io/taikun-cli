package application

import (
	"github.com/itera-io/taikun-cli/cmd/application/list"
	"github.com/spf13/cobra"
)

func NewCmdApplication() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "application <command>",
		Short:   "Explore applications in enabled repositories",
		Aliases: []string{"applications", "app", "apps"},
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
