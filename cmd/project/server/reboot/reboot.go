package reboot

import "github.com/spf13/cobra"

type RebootOptions struct {
	// FIXME add options
}

func NewCmdReboot() *cobra.Command {
	var opts RebootOptions

	cmd := cobra.Command{
		Use:   "reboot <command>",
		Short: "Manage reboots",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return rebootRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func rebootRun(opts *RebootOptions) (err error) {
	// FIXME
	return
}
