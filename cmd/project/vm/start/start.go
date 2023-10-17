package start

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type StartOptions struct {
	StandaloneVMID int32
}

func NewCmdStart() *cobra.Command {
	var opts StartOptions

	cmd := cobra.Command{
		Use:   "start <vm-id>",
		Short: "Start a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return startRun(&opts)
		},
	}

	return &cmd
}

func startRun(opts *StartOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.StartStandaloneVmCommand{
		Id: &opts.StandaloneVMID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsStart(context.TODO()).StartStandaloneVmCommand(body).Execute()
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

		body := models.StartStandaloneVMCommand{ID: opts.StandaloneVMID}
		params := stand_alone_actions.NewStandAloneActionsStartParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.StandAloneActions.StandAloneActionsStart(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
