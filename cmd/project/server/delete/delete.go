package delete

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	// FIXME add options
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <command>",
		Short: "Manage deletes",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	// FIXME

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	// FIXME
	return
}
