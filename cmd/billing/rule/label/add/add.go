package add

import "github.com/spf13/cobra"

func NewCmdAdd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "add <billing-rule-id>",
		Short: "Add a label to a billing rule",
	}

	return &cmd
}
