package unshelve

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return unshelveRun(cmd, &opts)
		},
	}

	return &cmd
}

func unshelveRun(cmd *cobra.Command, opts *UnshelveOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.UnshelveStandaloneVmCommand{
		Id: &opts.StandaloneVMID,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsUnshelve(ctx).UnshelveStandaloneVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
