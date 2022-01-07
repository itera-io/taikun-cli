package create

import "github.com/spf13/cobra"

type CreateOptions struct {
	// FIXME add options
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <command>",
		Short: "Manage creates",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return createRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func createRun(opts *CreateOptions) (err error) {
	// FIXME
	return
}
