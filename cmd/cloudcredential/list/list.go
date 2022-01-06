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
		Short: "List cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&config.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByFlag(cmd, &config.SortBy,
		models.AmazonCredentialsListDto{},
		models.OpenstackCredentialsListDto{},
		models.AzureCredentialsListDto{},
	)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	// TODO refactor by calling the respective methods of aws, azure, os
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

	response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
	if err != nil {
		return
	}
	credentialsAmazon := response.Payload.Amazon
	credentialsAmazonCount := int32(len(credentialsAmazon))

	credentialsAzure := response.Payload.Azure
	credentialsAzureCount := int32(len(credentialsAzure))

	credentialsOpenstack := response.Payload.Openstack
	credentialsOpenstackCount := int32(len(credentialsOpenstack))

	for credentialsAmazonCount < response.Payload.TotalCountAws {
		params = params.WithOffset(&credentialsAmazonCount)
		response, err = apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return
		}
		credentialsAmazon = append(credentialsAmazon, response.Payload.Amazon...)
		credentialsAmazonCount = int32(len(credentialsAmazon))
	}
	credentialsAmazonGeneric := make([]interface{}, len(credentialsAmazon))
	for i, credential := range credentialsAmazon {
		credentialsAmazonGeneric[i] = *credential
	}

	for credentialsAzureCount < response.Payload.TotalCountAzure {
		params = params.WithOffset(&credentialsAzureCount)
		response, err = apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return
		}
		credentialsAzure = append(credentialsAzure, response.Payload.Azure...)
		credentialsAzureCount = int32(len(credentialsAzure))
	}
	credentialsAzureGeneric := make([]interface{}, len(credentialsAzure))
	for i, credential := range credentialsAzure {
		credentialsAzureGeneric[i] = *credential
	}

	for credentialsOpenstackCount < response.Payload.TotalCountOpenstack {
		params = params.WithOffset(&credentialsOpenstackCount)
		response, err = apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return
		}
		credentialsOpenstack = append(credentialsOpenstack, response.Payload.Openstack...)
		credentialsOpenstackCount = int32(len(credentialsOpenstack))
	}
	credentialsOpenstackGeneric := make([]interface{}, len(credentialsOpenstack))
	for i, credential := range credentialsOpenstack {
		credentialsOpenstackGeneric[i] = *credential
	}

	format.PrintMultipleResults(
		[]interface{}{
			credentialsAmazonGeneric,
			credentialsAzureGeneric,
			credentialsOpenstackGeneric,
		},
		[]string{"AWS", "Azure", "OpenStack"},
		"id",
		"name",
		"organizationName",
		"createdBy",
		"isDefault",
		"isLocked",
	)

	return
}
