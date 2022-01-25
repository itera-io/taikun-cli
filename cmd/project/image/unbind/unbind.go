package unbind

import (
	"fmt"

	"github.com/spf13/cobra"
)

type UnbindOptions struct {
	// FIXME
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := cobra.Command{
		Use:   "unbind <image-bound-id>...",
		Short: "Unbind one or more images from a project",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// FIXME
			return unbindRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func unbindRun(opts *UnbindOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
