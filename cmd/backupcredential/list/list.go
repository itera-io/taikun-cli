package list

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit          int32
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List backup credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return fmt.Errorf("limit flag must be positive")
			}
			return listRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := s3_credentials.NewS3CredentialsListParams().WithV(cmdutils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}

	backupCredentials := []*models.BackupCredentialsListDto{}
	for {
		response, err := apiClient.Client.S3Credentials.S3CredentialsList(params, apiClient)
		if err != nil {
			return err
		}
		backupCredentials = append(backupCredentials, response.Payload.Data...)
		backupCredentialsCount := int32(len(backupCredentials))
		if opts.Limit != 0 && backupCredentialsCount >= opts.Limit {
			break
		}
		if backupCredentialsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&backupCredentialsCount)
	}

	if opts.Limit != 0 && int32(len(backupCredentials)) > opts.Limit {
		backupCredentials = backupCredentials[:opts.Limit]
	}

	cmdutils.PrettyPrint(backupCredentials)
	return
}
