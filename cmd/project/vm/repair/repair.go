package repair

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone"
	"github.com/itera-io/taikungoclient/models"
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.RepairStandAloneVMCommand{ProjectID: opts.ProjectID}
	params := stand_alone.NewStandAloneRepairParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAlone.StandAloneRepair(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
