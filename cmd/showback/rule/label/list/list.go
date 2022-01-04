package list

import "github.com/spf13/cobra"

type ListOptions struct {
	// TODO define options
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "list a showback rule's labels",
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
