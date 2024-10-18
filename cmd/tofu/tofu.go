package tofu

import (
	"github.com/itera-io/taikun-cli/cmd/tofu/project"
	"github.com/spf13/cobra"
)

func NewCmdTofu() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tofu <command>",
		Short: "Prepare/repair the project",
	}

	cmd.AddCommand(project.NewCmdVms())

	return cmd
}
