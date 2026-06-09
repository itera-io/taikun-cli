package transferownership

import (
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type TransferOwnershipOptions struct {
	UserID string
}

func NewCmdTransferOwnership() *cobra.Command {
	var opts TransferOwnershipOptions

	cmd := cobra.Command{
		Use:   "transfer-ownership <USER_ID>",
		Short: "Transfer account ownership",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.UserID = args[0]
			return transferOwnershipRun(cmd, &opts)
		},
	}

	return &cmd
}

func transferOwnershipRun(cmd *cobra.Command, opts *TransferOwnershipOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.TransferOwnershipCommand{}
	body.SetUserId(opts.UserID)

	response, err := myApiClient.Client.AccountsAPI.AccountsTransferOwnership(ctx).TransferOwnershipCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
