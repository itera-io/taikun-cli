package delete

import "github.com/spf13/cobra"

type DeleteOptions struct {
	// FIXME add options
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete some or all standalone VMs from a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return deleteRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	// FIXME
	return
}
