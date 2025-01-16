package shelve

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.ShelveStandAloneVmCommand{
		Id: &opts.StandaloneVMID,
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsShelve(context.TODO()).ShelveStandAloneVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
