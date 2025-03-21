package disable

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DisableOptions struct {
	StandaloneVMID int32
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <vm-id>",
		Short: "Disable an OpenStack standalone VM's public IP",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return disableRun(&opts)
		},
	}

	return &cmd
}

func disableRun(opts *DisableOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	enableIP := types.DisableVMPublicIP
	body := taikuncore.StandAloneVmIpManagementCommand{
		Id:   &opts.StandaloneVMID,
		Mode: *taikuncore.NewNullableString(&enableIP),
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.StandaloneAPI.StandaloneIpManagement(context.TODO()).StandAloneVmIpManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
