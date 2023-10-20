package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <backup-credential-id>...",
		Short: "Delete one or more backup credentials",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(backupCredentialID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.S3CredentialsAPI.S3credentialsDelete(context.TODO(), backupCredentialID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Backup credential", backupCredentialID)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := s3_credentials.NewS3CredentialsDeleteParams().WithV(taikungoclient.Version)
		params = params.WithID(backupCredentialID)

		_, _, err = apiClient.Client.S3Credentials.S3CredentialsDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Backup credential", backupCredentialID)
		}

		return
	*/
}
