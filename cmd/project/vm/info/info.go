package info

import "github.com/spf13/cobra"

type InfoOptions struct {
	// FIXME add options
}

func NewCmdInfo() *cobra.Command {
	var opts InfoOptions

	cmd := cobra.Command{
		Use:   "info <vm-id>",
		Short: "Get detailed information on a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return infoRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func infoRun(opts *InfoOptions) (err error) {
	// FIXME
	return
}
