package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/list"

	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List backup credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmdutils.AddLimitFlag(cmd)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := s3_credentials.NewS3CredentialsListParams().WithV(apiconfig.Version)
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
		if list.Limit != 0 && backupCredentialsCount >= list.Limit {
			break
		}
		if backupCredentialsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&backupCredentialsCount)
	}

	if list.Limit != 0 && int32(len(backupCredentials)) > list.Limit {
		backupCredentials = backupCredentials[:list.Limit]
	}

	format.PrintResults(backupCredentials,
		"id",
		"organizationName",
		"s3Name",
		"s3AccessKeyId",
		"s3Endpoint",
		"s3Region",
		"isDefault",
		"isLocked",
	)
	return
}
