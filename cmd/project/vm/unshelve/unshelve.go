package unshelve

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone_actions"
	"github.com/itera-io/taikungoclient/models"
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
				return cmderr.IDArgumentNotANumberError
			}
			return unshelveRun(&opts)
		},
	}

	return &cmd
}

func unshelveRun(opts *UnshelveOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.UnshelveStandaloneVMCommand{ID: opts.StandaloneVMID}
	params := stand_alone_actions.NewStandAloneActionsUnshelveParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneActions.StandAloneActionsUnshelve(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
