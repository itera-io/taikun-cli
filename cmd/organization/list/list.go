package list

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/config"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return format.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

	return cmd
}

func printResults(organizations []*models.OrganizationDetailsDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(organizations)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(organizations))
		for i, organization := range organizations {
			data[i] = organization
		}
		format.PrettyPrintTable(data,
			"id",
			"name",
			"fullName",
			"discountRate",
			"partnerName",
			"isEligibleUpdateSubscription",
			"isLocked",
			"isReadOnly",
			"users",
			"cloudCredentials",
			"projects",
			"servers",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := organizations.NewOrganizationsListParams().WithV(apiconfig.Version)
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var organizations = make([]*models.OrganizationDetailsDto, 0)
	for {
		response, err := apiClient.Client.Organizations.OrganizationsList(params, apiClient)
		if err != nil {
			return err
		}
		organizations = append(organizations, response.Payload.Data...)
		organizationsCount := int32(len(organizations))
		if opts.Limit != 0 && organizationsCount >= opts.Limit {
			break
		}
		if organizationsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&organizationsCount)
	}

	if opts.Limit != 0 && int32(len(organizations)) > opts.Limit {
		organizations = organizations[:opts.Limit]
	}

	printResults(organizations)
	return
}
