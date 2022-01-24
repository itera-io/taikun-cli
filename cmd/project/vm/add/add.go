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
		Use:   "add <project-id>",
		Short: "Add a standalone VM to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME
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
