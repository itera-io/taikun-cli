package unshelve

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UnshelveOptions struct {
	StandaloneVMID int32
}

func NewCmdUnshelve() *cobra.Command {
	var opts UnshelveOptions

	cmd := cobra.Command{
		Use:   "unshelve <vm-id>",
		Short: "Unshelve a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unshelveRun(&opts)
		},
	}

	return &cmd
}

func unshelveRun(opts *UnshelveOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.UnshelveStandaloneVmCommand{
		Id: &opts.StandaloneVMID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsUnshelve(context.TODO()).UnshelveStandaloneVmCommand(body).Execute()
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

		body := models.UnshelveStandaloneVMCommand{ID: opts.StandaloneVMID}
		params := stand_alone_actions.NewStandAloneActionsUnshelveParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.StandAloneActions.StandAloneActionsUnshelve(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
