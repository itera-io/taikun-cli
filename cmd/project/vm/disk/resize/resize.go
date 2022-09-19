package resize

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/stand_alone_vm_disks"
	"github.com/itera-io/taikungoclient/models"
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
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.UpdateStandaloneVMDiskSizeCommand{
		ID:   opts.DiskID,
		Size: opts.Size,
	}

	params := stand_alone_vm_disks.NewStandAloneVMDisksUpdateDiskSizeParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneVMDisks.StandAloneVMDisksUpdateDiskSize(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
