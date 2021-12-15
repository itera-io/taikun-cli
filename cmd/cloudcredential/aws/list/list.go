package list

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

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
				return fmt.Errorf("limit flag must be positive")
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

func ListRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(cmdutils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		cmdutils.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&cmdutils.SortDirection)
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

	cmdutils.PrettyPrint(amazonCloudCredentials)
	return
}
