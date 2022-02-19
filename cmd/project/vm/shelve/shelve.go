package shelve

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone_actions"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ShelveOptions struct {
	StandaloneVMID int32
}

func NewCmdShelve() *cobra.Command {
	var opts ShelveOptions

	cmd := cobra.Command{
		Use:   "shelve <vm-id>",
		Short: "Shelve a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return shelveRun(&opts)
		},
	}

	return &cmd
}

func shelveRun(opts *ShelveOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.ShelveStandAloneVMCommand{ID: opts.StandaloneVMID}
	params := stand_alone_actions.NewStandAloneActionsShelveParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneActions.StandAloneActionsShelve(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
