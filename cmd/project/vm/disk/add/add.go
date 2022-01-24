package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

type AddOptions struct {
	// FIXME add options
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <vm-id>",
		Short: "Add a disk to a standalone VM",
		Args:  cobra.ExactArgs(1), // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return addRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	fmt.Println("TODO")
	// FIXME
	return
}
