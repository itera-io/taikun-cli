package commit

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CommitOptions struct {
	// FIXME add options
}

func NewCmdCommit() *cobra.Command {
	var opts CommitOptions

	cmd := cobra.Command{
		Use:   "commit <project-id>",
		Short: "Commit changes to a project's standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return commitRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func commitRun(opts *CommitOptions) (err error) {
	fmt.Println("TODO")
	// FIXME
	return
}
