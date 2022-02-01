package start

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone_actions"
	"github.com/itera-io/taikungoclient/models"
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
				return cmderr.IDArgumentNotANumberError
			}
			return startRun(&opts)
		},
	}

	return &cmd
}

func startRun(opts *StartOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.StartStandaloneVMCommand{ID: opts.StandaloneVMID}
	params := stand_alone_actions.NewStandAloneActionsStartParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneActions.StandAloneActionsStart(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
