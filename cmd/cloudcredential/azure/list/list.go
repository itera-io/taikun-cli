package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

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
			return ListRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&config.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByFlag(cmd, models.AzureCredentialsListDto{})

	return cmd
}

func ListRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var azureCloudCredentials = make([]*models.AzureCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return err
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

	format.PrintResults(azureCloudCredentials,
		"id",
		"name",
		"organizationName",
		"location",
		"availabilityZone",
		"isDefault",
		"isLocked",
	)
	return
}
