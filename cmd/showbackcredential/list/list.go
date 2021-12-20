package list

import "github.com/spf13/cobra"

type ListOptions struct {
	// TODO set flags
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List showback credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	// TODO set flags

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	return
}
