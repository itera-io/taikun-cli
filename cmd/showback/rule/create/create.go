package create

import "github.com/spf13/cobra"

type CreateOptions struct {
	Name string
	// TODO add flags
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <name>",
		Short: "Create a showback rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return createRun(&opts)
		},
	}

	// TODO add flags

	return &cmd
}

func createRun(opts *CreateOptions) (err error) {
	// FIXME
	return
}
