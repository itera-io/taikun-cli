package disable

import "github.com/spf13/cobra"

type DisableOptions struct {
	// FIXME add options
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <command>", // FIXME
		Short: "Manage disables",   // FIXME
		Args:  cobra.NoArgs,        // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return disableRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func disableRun(opts *DisableOptions) (err error) {
	// FIXME
	return
}
