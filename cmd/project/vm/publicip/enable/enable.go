package enable

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return enableRun(&opts)
		},
	}

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	enableIP := types.EnableVMPublicIP
	body := taikuncore.StandAloneVmIpManagementCommand{
		Id:   &opts.StandaloneVMID,
		Mode: *taikuncore.NewNullableString(&enableIP),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneAPI.StandaloneIpManagement(context.TODO()).StandAloneVmIpManagementCommand(body).Execute()
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

		body := models.StandAloneVMIPManagementCommand{
			ID:   opts.StandaloneVMID,
			Mode: types.EnableVMPublicIP,
		}

		params := stand_alone.NewStandAloneIPManagementParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.StandAlone.StandAloneIPManagement(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
