package list

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/config"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/cloud_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List aws cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return format.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			return ListRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

	return cmd
}

func printResults(credentials []*models.AmazonCredentialsListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(credentials)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(credentials))
		for i, credential := range credentials {
			data[i] = credential
		}
		format.PrettyPrintTable(data,
			"id",
			"name",
			"organizationName",
			"region",
			"availabilityZone",
			"isDefault",
			"isLocked",
		)
	}
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
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
		fmt.Printf("sorting by %s\n", opts.SortBy)
	}

	var amazonCloudCredentials = make([]*models.AmazonCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err != nil {
			return err
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

	printResults(amazonCloudCredentials)
	return
}
