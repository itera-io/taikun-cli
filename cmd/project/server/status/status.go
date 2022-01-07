package status

import "github.com/spf13/cobra"

type StatusOptions struct {
	// FIXME add options
}

func NewCmdStatus() *cobra.Command {
	var opts StatusOptions

	cmd := cobra.Command{
		Use:   "status <command>",
		Short: "Manage statuss",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return statusRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func statusRun(opts *StatusOptions) (err error) {
	// FIXME
	return
}
