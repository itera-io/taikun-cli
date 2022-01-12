package delete

import (
	"log"

	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	KubeconfigIDs []int32
	// FIXME add options
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <kubeconfig-id>...",
		Short: "Delete one or more kubeconfigs",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return deleteRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	// FIXME
	log.Println("TODO")
	return
}
