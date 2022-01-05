package delete

import "github.com/spf13/cobra"

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <label-id>...",
		Short: "Delete one or more labels from a billing rule",
	}

	return &cmd
}
