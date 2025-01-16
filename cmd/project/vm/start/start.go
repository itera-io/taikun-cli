package start

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return startRun(&opts)
		},
	}

	return &cmd
}

func startRun(opts *StartOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.StartStandaloneVmCommand{
		Id: &opts.StandaloneVMID,
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsStart(context.TODO()).StartStandaloneVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
