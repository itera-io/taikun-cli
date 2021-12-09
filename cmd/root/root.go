package root

import (
	"taikun-cli/cmd/noop"
	"taikun-cli/cmd/user"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "taikun <command> <subcommand> [flags]",
		Short:        "Taikun CLI",
		Long:         `Manage Taikun resources from the command line.`,
		SilenceUsage: true,
	}

	cmd.AddCommand(noop.NewCmdNoop())
	cmd.AddCommand(user.NewCmdUser())

	return cmd
}
