package unshelve

import "github.com/spf13/cobra"

type UnshelveOptions struct {
	// FIXME add options
}

func NewCmdUnshelve() *cobra.Command {
	var opts UnshelveOptions

	cmd := cobra.Command{
		Use:   "unshelve <vm-id>",
		Short: "Unshelve a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return unshelveRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func unshelveRun(opts *UnshelveOptions) (err error) {
	// FIXME
	return
}
