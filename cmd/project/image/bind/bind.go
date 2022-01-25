package bind

import (
	"fmt"

	"github.com/spf13/cobra"
)

type BindOptions struct {
	// FIXME
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := cobra.Command{
		Use:   "bind <project-id>",
		Short: "Bind one or multiple images to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// FIXME
			return bindRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func bindRun(opts *BindOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
