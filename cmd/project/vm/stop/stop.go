package stop

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return stopRun(&opts)
		},
	}

	return &cmd
}

func stopRun(opts *StopOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.StopStandaloneVmCommand{
		Id: &opts.StandaloneVMID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsStop(context.TODO()).StopStandaloneVmCommand(body).Execute()
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

		body := models.StopStandaloneVMCommand{ID: opts.StandaloneVMID}
		params := stand_alone_actions.NewStandAloneActionsStopParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.StandAloneActions.StandAloneActionsStop(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
