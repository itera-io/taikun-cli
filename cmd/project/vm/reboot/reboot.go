package reboot

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/stand_alone_actions"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type RebootOptions struct {
	StandaloneVMID int32
	HardReboot     bool
}

func NewCmdReboot() *cobra.Command {
	var opts RebootOptions

	cmd := cobra.Command{
		Use:   "reboot <vm-id>",
		Short: "Reboot a standalone VM",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return rebootRun(&opts)
		},
	}

	cmd.Flags().BoolVar(&opts.HardReboot, "hard", false, "Hard reboot")

	return &cmd
}

func rebootRun(opts *RebootOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.RebootStandAloneVMCommand{
		ID:   opts.StandaloneVMID,
		Type: types.GetVMRebootType(opts.HardReboot),
	}

	params := stand_alone_actions.NewStandAloneActionsRebootParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAloneActions.StandAloneActionsReboot(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
