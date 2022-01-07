package list

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	// FIXME add options
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <command>",
		Short: "Manage lists",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	// FIXME

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	return
}
