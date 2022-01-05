package list

import "github.com/spf13/cobra"

func NewCmdList() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list <billing-rule-id>",
		Short: "List a billing rule's labels",
		Args:  cobra.ExactArgs(1),
	}

	return &cmd
}
