package list

import "github.com/spf13/cobra"

type ListOptions struct {
	// FIXME add options
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <vm-id>",
		Short: "List a standalone VM's disks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return listRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	return
}
