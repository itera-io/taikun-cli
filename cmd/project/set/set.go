package set

import (
	"github.com/itera-io/taikun-cli/cmd/project/set/expiration"
	"github.com/spf13/cobra"
)

func NewCmdSet() *cobra.Command {
	cmd := cobra.Command{
		Use:   "set <command>",
		Short: "Manage a project's life time",
	}

	cmd.AddCommand(expiration.NewCmdExpiration())

	return &cmd
}
