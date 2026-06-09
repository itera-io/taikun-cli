package enable

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return enableRun(cmd, &opts)
		},
	}

	return &cmd
}

func enableRun(cmd *cobra.Command, opts *EnableOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	enableIP := types.EnableVMPublicIP
	body := taikuncore.StandAloneVmIpManagementCommand{
		Id:   &opts.StandaloneVMID,
		Mode: *taikuncore.NewNullableString(&enableIP),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneAPI.StandaloneIpManagement(ctx).StandAloneVmIpManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
