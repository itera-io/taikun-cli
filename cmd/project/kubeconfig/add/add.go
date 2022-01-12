package add

import (
	"log"

	"github.com/spf13/cobra"
)

type AddOptions struct {
	ProjectID int32
	// FIXME add options
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <project-id>",
		Short: "Add a kubeconfig",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return addRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// FIXME
	log.Println("TODO")
	return
}
