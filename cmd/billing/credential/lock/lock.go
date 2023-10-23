package lock

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <billing-credential-id>",
		Short: "Lock a billing credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(id)
		},
	}

	return cmd
}

func lockRun(billingCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.OperationCredentialLockManagerCommand{
		Id:   &billingCredentialID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.OperationCredentialsAPI.OpscredentialsLockManager(context.TODO()).OperationCredentialLockManagerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.OperationCredentialLockManagerCommand{
			ID:   billingCredentialID,
			Mode: types.LockedMode,
		}
		params := ops_credentials.NewOpsCredentialsLockManagerParams().WithV(taikungoclient.Version).WithBody(&body)

		_, err = apiClient.Client.OpsCredentials.OpsCredentialsLockManager(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
