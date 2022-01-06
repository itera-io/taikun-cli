package list

import (
	"log"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
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

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByFlag(cmd, &opts.SortBy,
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

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	log.Println("before first request")

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

	log.Println("after first request")

	for credentialsAmazonCount < response.Payload.TotalCountAws {
		params = params.WithOffset(&credentialsAmazonCount)
		response, err = apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return
		}
		credentialsAmazon = append(credentialsAmazon, response.Payload.Amazon...)
		credentialsAmazonCount = int32(len(credentialsAmazon))
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

	for credentialsOpenstackCount < response.Payload.TotalCountOpenstack {
		params = params.WithOffset(&credentialsOpenstackCount)
		response, err = apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return
		}
		credentialsOpenstack = append(credentialsOpenstack, response.Payload.Openstack...)
		credentialsOpenstackCount = int32(len(credentialsOpenstack))
	}

	log.Println("printing multiple results")
	format.PrintMultipleResults(
		[]interface{}{
			credentialsAmazon,
			credentialsAzure,
			credentialsOpenstack,
		},
		cmdutils.GetCommonJsonTagsInStructs(
			[]interface{}{
				models.AmazonCredentialsListDto{},
				models.OpenstackCredentialsListDto{},
				models.AzureCredentialsListDto{},
			},
		)...,
	)

	return
}
