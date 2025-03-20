package remove

import (
	"context"
	"errors"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/project/vm/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	DeleteAll bool
	ProjectID int32
	VMIDs     []int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete some or all standalone VMs from a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			if opts.DeleteAll {
				if len(opts.VMIDs) != 0 {
					return errors.New("Cannot set both --vm-ids and --all-project flags")
				}
			} else {
				if len(opts.VMIDs) == 0 {
					return errors.New("Must set one of --vm-ids and --all-project flags")
				}
			}
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32SliceVarP(&opts.VMIDs, "vm-ids", "v", []int32{}, "IDs of the standalone VMs to delete")
	cmd.Flags().BoolVarP(&opts.DeleteAll, "all-project", "a", false, "Delete all of the project's standalone VMs")

	return &cmd
}

func deleteRun(opts *DeleteOptions) error {
	myApiClient := tk.NewClient()
	body := taikuncore.ProjectDeploymentDeleteVmsCommand{
		ProjectId: &opts.ProjectID,
	}

	if opts.DeleteAll {
		allVMs, err := list.ListVMs(&list.ListOptions{ProjectID: opts.ProjectID})
		if err != nil {
			return err
		}

		if len(allVMs) == 0 {
			return fmt.Errorf("project %d has no standalone VMs", opts.ProjectID)
		}

		allVMIDs := make([]int32, len(allVMs))
		for i, vm := range allVMs {
			allVMIDs[i] = vm.GetId()
		}
		body.VmIds = allVMIDs
	} else {
		body.VmIds = opts.VMIDs
	}

	_, response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDeleteVms(context.TODO()).ProjectDeploymentDeleteVmsCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	for _, id := range body.VmIds {
		out.PrintDeleteSuccess("Standalone VM", id)
	}

	return nil

}
