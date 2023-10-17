package status

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type StatusOptions struct {
	StandaloneVMID int32
}

func NewCmdStatus() *cobra.Command {
	var opts StatusOptions

	cmd := cobra.Command{
		Use:   "status <vm-id>",
		Short: "Show a standalone VM's status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return statusRun(&opts)
		},
	}

	return &cmd
}

func statusRun(opts *StatusOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsStatus(context.TODO(), opts.StandaloneVMID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	out.Println(data)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := stand_alone_actions.NewStandAloneActionsShowStandaloneVMStatusParams().WithV(taikungoclient.Version)
		params = params.WithID(opts.StandaloneVMID)

		response, err := apiClient.Client.StandAloneActions.StandAloneActionsShowStandaloneVMStatus(params, apiClient)
		if err == nil {
			out.Println(response.Payload)
		}

		return
	*/
}
