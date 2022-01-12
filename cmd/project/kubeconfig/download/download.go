package download

import (
	"fmt"

	"github.com/spf13/cobra"
)

type DownloadOptions struct {
	KubeconfigID int32
}

func NewCmdDownload() *cobra.Command {
	var opts DownloadOptions

	cmd := cobra.Command{
		Use:   "download <kubeconfig-id>",
		Short: "Download a kubeconfig",
		Args:  cobra.NoArgs, // FIXME maybe
		RunE: func(cmd *cobra.Command, args []string) error {
			// FIXME maybe
			return downloadRun(&opts)
		},
	}

	// FIXME

	return &cmd
}

func downloadRun(opts *DownloadOptions) (err error) {
	// FIXME
	fmt.Println("TODO")
	return
}
