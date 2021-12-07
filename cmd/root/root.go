package root

import (
	"taikun-cli/cmd/noop"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taikun <command> <subcommand> [flags]",
		Short: "Taikun CLI",
		Long:  `Manage Taikun resources from the command line.`,
	}

	cmd.AddCommand(noop.NewCmdNoop())

	return cmd
}
