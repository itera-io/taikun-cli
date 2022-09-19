package remove

import (
	"errors"
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/project/vm/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/stand_alone"
	"github.com/itera-io/taikungoclient/models"
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
					return errors.New("Cannot set both --vm-ids and --all-vms flags")
				}
			} else {
				if len(opts.VMIDs) == 0 {
					return errors.New("Must set one of --vm-ids and --all-vms flags")
				}
			}
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32SliceVarP(&opts.VMIDs, "vm-ids", "v", []int32{}, "IDs of the standalone VMs to delete")
	cmd.Flags().BoolVarP(&opts.DeleteAll, "all-vms", "a", false, "Delete all of the project's standalone VMs")

	return &cmd
}

func deleteRun(opts *DeleteOptions) error {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return err
	}

	body := models.DeleteStandAloneVMCommand{
		ProjectID: opts.ProjectID,
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
			allVMIDs[i] = vm.ID
		}

		body.VMIds = allVMIDs
	} else {
		body.VMIds = opts.VMIDs
	}

	params := stand_alone.NewStandAloneDeleteParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	if _, err = apiClient.Client.StandAlone.StandAloneDelete(params, apiClient); err != nil {
		return err
	}

	for _, id := range body.VMIds {
		out.PrintDeleteSuccess("Standalone VM", id)
	}

	return nil
}
