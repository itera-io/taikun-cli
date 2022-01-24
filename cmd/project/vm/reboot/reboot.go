package reboot

import (
	"fmt"

	"github.com/spf13/cobra"
)

type RebootOptions struct {
	// FIXME add options
}

func NewCmdReboot() *cobra.Command {
	var opts RebootOptions

	cmd := cobra.Command{
		Use:   "reboot <vm-id>",
		Short: "Reboot a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return rebootRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func rebootRun(opts *RebootOptions) (err error) {
	fmt.Println("TODO")
	// FIXME
	return
}
