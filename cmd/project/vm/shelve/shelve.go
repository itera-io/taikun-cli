package shelve

import "github.com/spf13/cobra"

type ShelveOptions struct {
	// FIXME add options
}

func NewCmdShelve() *cobra.Command {
	var opts ShelveOptions

	cmd := cobra.Command{
		Use:   "shelve <vm-id>",
		Short: "Shelve a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return shelveRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func shelveRun(opts *ShelveOptions) (err error) {
	// FIXME
	return
}
