package stop

import "github.com/spf13/cobra"

type StopOptions struct {
	// FIXME add options
}

func NewCmdStop() *cobra.Command {
	var opts StopOptions

	cmd := cobra.Command{
		Use:   "stop <vm-id>",
		Short: "Stop a VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return stopRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func stopRun(opts *StopOptions) (err error) {
	// FIXME
	return
}
