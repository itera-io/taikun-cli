package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	// FIXME add options
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return listRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	fmt.Println("TODO")
	// FIXME
	return
}
