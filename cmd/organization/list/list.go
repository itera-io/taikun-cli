package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")

	cmdutils.AddSortByFlag(cmd, &opts.SortBy, models.OrganizationDetailsDto{})
	cmdutils.AddLimitFlag(cmd)

	return cmd
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
		if config.Limit != 0 && organizationsCount >= config.Limit {
			break
		}
		if organizationsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&organizationsCount)
	}

	if config.Limit != 0 && int32(len(organizations)) > config.Limit {
		organizations = organizations[:config.Limit]
	}

	format.PrintResults(organizations,
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
	return
}
