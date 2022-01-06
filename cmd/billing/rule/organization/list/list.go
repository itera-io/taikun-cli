package list

import "github.com/spf13/cobra"

type ListOptions struct {
	// FIXME
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <billing-rule-id>",
		Short: "List a billing rule's bound organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME
			return listRun(&opts)
		},
	}

	// FIXME add flags

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	return
}
