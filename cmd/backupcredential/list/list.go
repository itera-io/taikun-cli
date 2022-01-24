package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"S3-NAME", "s3Name",
		),
		field.NewVisible(
			"S3-ACCESS-KEY-ID", "s3AccessKeyId",
		),
		field.NewHidden(
			"S3-ENDPOINT", "s3Endpoint",
		),
		field.NewVisible(
			"S3-REGION", "s3Region",
		),
		field.NewVisible(
			"DEFAULT", "isDefault",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List backup credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := s3_credentials.NewS3CredentialsListParams().WithV(api.Version)
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

	out.PrintResults(backupCredentials, listFields)
	return
}
