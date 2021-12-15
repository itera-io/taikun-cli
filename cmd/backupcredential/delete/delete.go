package delete

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <backup-credential-id>",
		Short: "Delete a backup credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			backupCredentialID, err := utils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given ID must be a number")
			}
			return deleteRun(backupCredentialID)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := s3_credentials.NewS3CredentialsDeleteParams().WithV(utils.ApiVersion)
	params = params.WithID(id)
	_, _, err = apiClient.Client.S3Credentials.S3CredentialsDelete(params, apiClient)
	if err == nil {
		utils.PrintDeleteSuccess("Backup credential", id)
	}

	return
}
