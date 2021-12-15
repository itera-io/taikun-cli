package delete

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <cloud-credential-id>",
		Short: "Delete a cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := utils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given id must be a number")
			}
			return deleteRun(id)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsDeleteParams().WithV(utils.ApiVersion).WithCloudID(id)
	_, _, err = apiClient.Client.CloudCredentials.CloudCredentialsDelete(params, apiClient)
	if err == nil {
		utils.PrintDeleteSuccess("Cloud credential", id)
	}

	return
}
