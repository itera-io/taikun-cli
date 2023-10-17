package lock

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <cloud-credential-id>",
		Short: "Lock a cloud credential",
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

func lockRun(cloudCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CloudLockManagerCommand{
		Id:   &cloudCredentialID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.CloudCredentialAPI.CloudcredentialsLockManager(context.TODO()).CloudLockManagerCommand(body).Execute()
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

		body := &models.CloudLockManagerCommand{
			ID:   cloudCredentialID,
			Mode: types.LockedMode,
		}
		params := cloud_credentials.NewCloudCredentialsLockManagerParams().WithV(taikungoclient.Version).WithBody(body)

		_, err = apiClient.Client.CloudCredentials.CloudCredentialsLockManager(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
