package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"BILLING-ACCOUNT", "billingAccountName",
		),
		field.NewVisible(
			"FOLDER-ID", "folderId",
		),
		field.NewVisible(
			"PARTNER", "partnerName",
		),
		field.NewVisible(
			"REGION", "region",
		),
		field.NewVisible(
			"ZONE", "zone",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewVisibleWithToStringFunc(
			"ISLOCKED", "isLocked", out.FormatLockStatus,
		),
		// TODO add fields
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
		Short: "List Google Cloud Platform credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "cloud-credentials", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	googleCloudCredentials, err := ListCloudCredentialsGoogle(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(googleCloudCredentials, listFields)
}

func ListCloudCredentialsGoogle(opts *ListOptions) (credentials []interface{}, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}

	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var googleCloudCredentials = make([]*models.GoogleCredentialsListDto, 0)

	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return nil, err
		}

		googleCloudCredentials = append(googleCloudCredentials, response.Payload.Google...)

		count := int32(len(googleCloudCredentials))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == response.Payload.TotalCountAws {
			break
		}

		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(googleCloudCredentials)) > opts.Limit {
		googleCloudCredentials = googleCloudCredentials[:opts.Limit]
	}

	credentials = make([]interface{}, len(googleCloudCredentials))
	for i, credential := range googleCloudCredentials {
		credentials[i] = *credential
	}

	return credentials, nil
}
