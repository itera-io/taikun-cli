package resize

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type ResizeOptions struct {
	DiskID int32
	Size   int64
}

func NewCmdResize() *cobra.Command {
	var opts ResizeOptions

	cmd := cobra.Command{
		Use:   "resize <disk-id>",
		Short: "Resize a standalone VM's disk",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.DiskID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return resizeRun(&opts)
		},
	}

	cmd.Flags().Int64VarP(&opts.Size, "size", "s", 0, "New size in GiB (required)")
	cmdutils.MarkFlagRequired(&cmd, "size")

	return &cmd
}

func resizeRun(opts *ResizeOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.UpdateStandaloneVmDiskSizeCommand{
		Id:   &opts.DiskID,
		Size: &opts.Size,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneVMDisksAPI.StandalonevmdisksUpdateSize(context.TODO()).UpdateStandaloneVmDiskSizeCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
