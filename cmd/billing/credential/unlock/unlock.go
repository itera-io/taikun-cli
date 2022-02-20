package unlock

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/ops_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <billing-credential-id>",
		Short: "Unlock a billing credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unlockRun(id)
		},
	}

	return cmd
}

func unlockRun(billingCredentialID int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.OperationCredentialLockManagerCommand{
		ID:   billingCredentialID,
		Mode: types.UnlockedMode,
	}
	params := ops_credentials.NewOpsCredentialsLockManagerParams().WithV(api.Version).WithBody(&body)

	_, err = apiClient.Client.OpsCredentials.OpsCredentialsLockManager(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
