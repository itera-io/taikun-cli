package start

import "github.com/spf13/cobra"

type StartOptions struct {
	// FIXME add options
}

func NewCmdStart() *cobra.Command {
	var opts StartOptions

	cmd := cobra.Command{
		Use:   "start <vm-id>",
		Short: "Start a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return startRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func startRun(opts *StartOptions) (err error) {
	// FIXME
	return
}
