package rename

import "github.com/spf13/cobra"

type RenameOptions struct {
	// FIXME add options
}

func NewCmdRename() *cobra.Command {
	var opts RenameOptions

	cmd := cobra.Command{
		Use:   "rename <standalone-profile-id>",
		Short: "Rename a standalone profile",
		Args:  cobra.ExactArgs(1), // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return renameRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func renameRun(opts *RenameOptions) (err error) {
	// FIXME
	return
}
