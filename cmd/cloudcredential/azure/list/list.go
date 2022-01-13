package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
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
		Short: "List Azure cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByAndReverseFlags(cmd, models.AzureCredentialsListDto{})

	return cmd
}

func listRun(opts *ListOptions) error {
	azureCloudCredentials, err := ListCloudCredentialsAzure(opts)
	if err != nil {
		return err
	}

	out.PrintResults(azureCloudCredentials,
		"id",
		"name",
		"organizationName",
		"location",
		"availabilityZone",
		"isDefault",
		"isLocked",
	)

	return nil
}

func ListCloudCredentialsAzure(opts *ListOptions) (credentials []interface{}, err error) {
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

	var azureCloudCredentials = make([]*models.AzureCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return nil, err
		}
		azureCloudCredentials = append(azureCloudCredentials, response.Payload.Azure...)
		count := int32(len(azureCloudCredentials))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCountAzure {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(azureCloudCredentials)) > config.Limit {
		azureCloudCredentials = azureCloudCredentials[:config.Limit]
	}

	credentials = make([]interface{}, len(azureCloudCredentials))
	for i, credential := range azureCloudCredentials {
		credentials[i] = *credential
	}

	return
}
