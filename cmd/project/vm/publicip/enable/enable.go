package enable

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/stand_alone"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	StandaloneVMID int32
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <vm-id>",
		Short: "Enable an OpenStack standalone VM's public IP",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return enableRun(&opts)
		},
	}

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.StandAloneVMIPManagementCommand{
		ID:   opts.StandaloneVMID,
		Mode: types.EnableVMPublicIP,
	}

	params := stand_alone.NewStandAloneIPManagementParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.StandAlone.StandAloneIPManagement(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
