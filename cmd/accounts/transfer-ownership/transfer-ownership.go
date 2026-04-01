package transferownership

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
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
			return transferOwnershipRun(&opts)
		},
	}

	return &cmd
}

func transferOwnershipRun(opts *TransferOwnershipOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.TransferOwnershipCommand{}
	body.SetUserId(opts.UserID)

	response, err := myApiClient.Client.AccountsAPI.AccountsTransferOwnership(context.TODO()).TransferOwnershipCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
