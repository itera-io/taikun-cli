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
		Use:   "delete <standalone-profile-id>...",
		Short: "Delete one or more standalone profiles",
		Args:  cobra.MinimumNArgs(1),
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
