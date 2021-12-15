package lock

import (
	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <cloud-credential-id>",
		Short: "Lock a cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := utils.Atoi32(args[0])
			if err != nil {
				return utils.WrongIDArgumentFormatError
			}
			return lockRun(id)
		},
	}

	return cmd
}

func lockRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CloudLockManagerCommand{
		ID:   id,
		Mode: "lock",
	}
	params := cloud_credentials.NewCloudCredentialsLockManagerParams().WithV(utils.ApiVersion).WithBody(body)
	_, err = apiClient.Client.CloudCredentials.CloudCredentialsLockManager(params, apiClient)
	if err == nil {
		utils.PrintStandardSuccess()
	}

	return
}
