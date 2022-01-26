package stop

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone_actions"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type StopOptions struct {
	StandaloneVMID int32
}

func NewCmdStop() *cobra.Command {
	var opts StopOptions

	cmd := cobra.Command{
		Use:   "stop <vm-id>",
		Short: "Stop a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return stopRun(&opts)
		},
	}

	return &cmd
}

func stopRun(opts *StopOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.StopStandaloneVMCommand{ID: opts.StandaloneVMID}
	params := stand_alone_actions.NewStandAloneActionsStopParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneActions.StandAloneActionsStop(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
