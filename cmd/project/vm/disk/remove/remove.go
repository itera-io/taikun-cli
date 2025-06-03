package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.DeleteVmDiskCommand{
		StandaloneVmId: &opts.StandaloneVMID,
		VmDiskIds:      opts.DiskIDs,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDeleteVmDisks(context.TODO()).DeleteVmDiskCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	for _, id := range body.VmDiskIds {
		out.PrintDeleteSuccess("Standalone VM disk", id)
	}
	return

}
