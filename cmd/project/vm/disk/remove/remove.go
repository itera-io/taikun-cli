package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/stand_alone_vm_disks"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	StandaloneVMID int32
	DiskIDs        []int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <vm-id>",
		Short: "Delete one or more disks from a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32SliceVarP(&opts.DiskIDs, "disk-ids", "d", []int32{}, "IDs of the disks to delete (required)")
	cmdutils.MarkFlagRequired(&cmd, "disk-ids")

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteStandAloneVMDiskCommand{
		StandaloneVMID: opts.StandaloneVMID,
		VMDiskIds:      opts.DiskIDs,
	}

	params := stand_alone_vm_disks.NewStandAloneVMDisksDeleteParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneVMDisks.StandAloneVMDisksDelete(params, apiClient)
	if err == nil {
		for _, id := range body.VMDiskIds {
			out.PrintDeleteSuccess("Standalone VM disk", id)
		}
	}

	return
}
