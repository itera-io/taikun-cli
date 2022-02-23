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
		Use:   "check <google-credential-filename>",
		Short: "Check the validity of a Google Cloud Platform credential",
		Args:  cobra.ExactArgs(1),
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
