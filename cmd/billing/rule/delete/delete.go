package delete

import "github.com/spf13/cobra"

type DeleteOptions struct {
	// TODO add options
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <id>...",
		Short: "Delete one or more billing rules",
		// TODO define Args
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteRun(&opts)
		},
	}

	// TODO set flags

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	// FIXME
	return
}
