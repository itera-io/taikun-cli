package enable

import "github.com/spf13/cobra"

type EnableOptions struct {
	// FIXME add options
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <command>", // FIXME
		Short: "Manage enables",   // FIXME
		Args:  cobra.NoArgs,       // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return enableRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	// FIXME
	return
}
