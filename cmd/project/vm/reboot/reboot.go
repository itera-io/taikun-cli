package reboot

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type RebootOptions struct {
	StandaloneVMID int32
	RebootType     bool
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
			return rebootRun(cmd, &opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.RebootType, "hard", "f", false, "Force hard reboot of server")

	return &cmd
}

func rebootRun(cmd *cobra.Command, opts *RebootOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.RebootStandAloneVmCommand{}
	body.SetId(opts.StandaloneVMID)
	if opts.RebootType {
		body.SetType("HARD")
	} else {
		body.SetType("SOFT")
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.StandaloneActionsAPI.StandaloneactionsReboot(ctx).RebootStandAloneVmCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
