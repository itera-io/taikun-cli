package check

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CheckOptions struct {
	// FIXME
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := cobra.Command{
		Use:   "check <command>",
		Short: "Manage checks", // FIXME
		Args:  cobra.NoArgs,    // FIXME
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// FIXME
			return checkRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func checkRun(opts *CheckOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
