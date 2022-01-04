package add

import "github.com/spf13/cobra"

type AddOptions struct {
	// TODO define options
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add",
		Short: "Add a label to a showback rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			return addRun(&opts)
		},
	}

	// TODO set flags

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// FIXME
	return
}
