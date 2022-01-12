package list

import (
	"log"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's kubeconfigs",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return listRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// FIXME
	log.Println("TODO")
	return
}
