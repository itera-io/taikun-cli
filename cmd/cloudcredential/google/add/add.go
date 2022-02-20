package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

type AddOptions struct {
	// FIXME
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <command>",
		Short: "Manage adds", // FIXME
		Args:  cobra.NoArgs,  // FIXME
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// FIXME
			return addRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
