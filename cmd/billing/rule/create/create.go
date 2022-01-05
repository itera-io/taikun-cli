package create

import "github.com/spf13/cobra"

type CreateOptions struct {
	// TODO add options
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <name>",
		Short: "Create a billing rule",
		// TODO define Args
		RunE: func(cmd *cobra.Command, args []string) error {
			return createRun(&opts)
		},
	}

	// TODO set flags

	return &cmd
}

func createRun(opts *CreateOptions) (err error) {
	// FIXME
	return
}
