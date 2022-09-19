package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient"
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
			"PROJECT", "project",
		),
		field.NewVisible(
			"USER", "user",
		),
		field.NewHidden(
			"DOMAIN", "domain",
		),
		field.NewHidden(
			"IMPORT-NETWORK", "importNetwork",
		),
		field.NewHidden(
			"PUBLIC-NETWORK", "publicNetwork",
		),
		field.NewHidden(
			"REGION", "region",
		),
		field.NewHidden(
			"TENANT-ID", "tenantId",
		),
		field.NewHidden(
			"URL", "url",
		),
		field.NewVisible(
			"DEFAULT", "isDefault",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewVisible(
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
		Short: "List OpenStack cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "cloud-credentials", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) error {
	openstackCloudCredentials, err := ListCloudCredentialsOpenStack(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(openstackCloudCredentials, listFields)
}

func ListCloudCredentialsOpenStack(opts *ListOptions) (credentials []interface{}, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return nil, err
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(taikungoclient.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}

	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var openstackCloudCredentials = make([]*models.OpenstackCredentialsListDto, 0)

	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return nil, err
		}

		openstackCloudCredentials = append(openstackCloudCredentials, response.Payload.Openstack...)

		count := int32(len(openstackCloudCredentials))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == response.Payload.TotalCountOpenstack {
			break
		}

		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(openstackCloudCredentials)) > opts.Limit {
		openstackCloudCredentials = openstackCloudCredentials[:opts.Limit]
	}

	credentials = make([]interface{}, len(openstackCloudCredentials))
	for i, credential := range openstackCloudCredentials {
		credentials[i] = *credential
	}

	return credentials, nil
}
