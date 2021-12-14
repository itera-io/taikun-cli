package unlock

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdUnlock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unlock <cloud-credential-id>",
		Short: "Unlock a cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given id must be a number")
			}
			return unlockRun(id)
		},
	}

	return cmd
}

func unlockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CloudLockManagerCommand{
		ID:   id,
		Mode: "unlock",
	}
	params := cloud_credentials.NewCloudCredentialsLockManagerParams().WithV(cmdutils.ApiVersion).WithBody(body)
	_, err = apiClient.Client.CloudCredentials.CloudCredentialsLockManager(params, apiClient)
	if err == nil {
		fmt.Println("Cloud Credential unlocked")
	}

	return
}
