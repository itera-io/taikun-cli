package status

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type StatusOptions struct {
	StandaloneVMID int32
}

func NewCmdStatus() *cobra.Command {
	var opts StatusOptions

	cmd := cobra.Command{
		Use:   "status <vm-id>",
		Short: "Show a standalone VM's status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.StandaloneVMID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return statusRun(cmd, &opts)
		},
	}

	return &cmd
}

func statusRun(cmd *cobra.Command, opts *StatusOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsStatus(ctx, opts.StandaloneVMID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	out.Println(data)
	return

}
