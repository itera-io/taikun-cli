package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	// FIXME
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <command>",
		Short: "Manage lists", // FIXME
		Args:  cobra.NoArgs,   // FIXME
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// FIXME
			return listRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
