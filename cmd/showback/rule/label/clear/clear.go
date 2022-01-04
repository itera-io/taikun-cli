package clear

import "github.com/spf13/cobra"

type ClearOptions struct {
	// TODO define options
}

func NewCmdClear() *cobra.Command {
	var opts ClearOptions

	cmd := cobra.Command{
		Use:   "clear",
		Short: "clear a showback rule's labels",
		RunE: func(cmd *cobra.Command, args []string) error {
			return clearRun(&opts)
		},
	}

	// TODO set flags

	return &cmd
}

func clearRun(opts *ClearOptions) (err error) {
	// FIXME
	return
}
