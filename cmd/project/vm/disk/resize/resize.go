package resize

import "github.com/spf13/cobra"

type ResizeOptions struct {
	// FIXME add options
}

func NewCmdResize() *cobra.Command {
	var opts ResizeOptions

	cmd := cobra.Command{
		Use:   "resize <disk-id>",
		Short: "Resize a standalone VM's disk",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return resizeRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func resizeRun(opts *ResizeOptions) (err error) {
	// FIXME
	return
}
