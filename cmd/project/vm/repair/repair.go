package repair

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type RepairOptions struct {
	ProjectID int32
}

func NewCmdRepair() *cobra.Command {
	var opts RepairOptions

	cmd := cobra.Command{
		Use:   "repair <project-id>",
		Short: "Repair a project's standalone VMs",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return repairRun(&opts)
		},
	}

	return &cmd
}

func repairRun(opts *RepairOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.RepairStandAloneVmCommand{
		ProjectId: &opts.ProjectID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneAPI.StandaloneRepair(context.TODO()).RepairStandAloneVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.RepairStandAloneVMCommand{ProjectID: opts.ProjectID}
		params := stand_alone.NewStandAloneRepairParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.StandAlone.StandAloneRepair(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
