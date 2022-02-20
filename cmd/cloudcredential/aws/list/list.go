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
			"REGION", "region",
		),
		field.NewVisible(
			"AVAILABILITY-ZONE", "availabilityZone",
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

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List AWS cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(cmd, "cloud-credentials", listFields)
	cmdutils.AddColumnsFlag(cmd, listFields)

	return cmd
}

func listRun(opts *ListOptions) error {
	amazonCloudCredentials, err := ListCloudCredentialsAws(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(amazonCloudCredentials, listFields)
}

func ListCloudCredentialsAws(opts *ListOptions) (credentials []interface{}, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}

	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var amazonCloudCredentials = make([]*models.AmazonCredentialsListDto, 0)

	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return nil, err
		}

		amazonCloudCredentials = append(amazonCloudCredentials, response.Payload.Amazon...)

		count := int32(len(amazonCloudCredentials))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == response.Payload.TotalCountAws {
			break
		}

		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(amazonCloudCredentials)) > opts.Limit {
		amazonCloudCredentials = amazonCloudCredentials[:opts.Limit]
	}

	credentials = make([]interface{}, len(amazonCloudCredentials))
	for i, credential := range amazonCloudCredentials {
		credentials[i] = *credential
	}

	return credentials, nil
}
