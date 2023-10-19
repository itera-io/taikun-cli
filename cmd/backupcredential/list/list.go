package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	cmdutils.AddSortByAndReverseFlags(&cmd, "backupcredentials", listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.S3CredentialsAPI.S3credentialsList(context.TODO())
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	var backupCredentials = make([]taikuncore.BackupCredentialsListDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		backupCredentials = append(backupCredentials, data.GetData()...)
		backupCredentialsCount := int32(len(backupCredentials))

		if opts.Limit != 0 && backupCredentialsCount >= opts.Limit {
			break
		}

		if backupCredentialsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(backupCredentialsCount)
	}

	if opts.Limit != 0 && int32(len(backupCredentials)) > opts.Limit {
		backupCredentials = backupCredentials[:opts.Limit]
	}

	return out.PrintResults(backupCredentials, listFields)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := s3_credentials.NewS3CredentialsListParams().WithV(taikungoclient.Version)
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

		return out.PrintResults(backupCredentials, listFields)
	*/
}
