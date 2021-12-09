package root

import (
	"taikun-cli/api"
	"taikun-cli/cmd/noop"
	"taikun-cli/cmd/user"

	"github.com/spf13/cobra"
)

func NewCmdRoot(apiClient *api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "taikun <command> <subcommand> [flags]",
		Short: "Taikun CLI",
		Long:  `Manage Taikun resources from the command line.`,
	}

	cmd.AddCommand(noop.NewCmdNoop(apiClient))
	cmd.AddCommand(user.NewCmdUser(apiClient))

	return cmd
}
