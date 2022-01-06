package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikungoclient/client/showback"
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

	cmd := cobra.Command{
		Use:   "list",
		Short: "List showback rules",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddSortByFlag(&cmd, &opts.SortBy, models.AccessProfilesListDto{})

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := showback.NewShowbackRulesListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var showbackRules = make([]*models.ShowbackRulesListDto, 0)
	for {
		response, err := apiClient.Client.Showback.ShowbackRulesList(params, apiClient)
		if err != nil {
			return err
		}
		showbackRules = append(showbackRules, response.Payload.Data...)
		count := int32(len(showbackRules))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(showbackRules)) > config.Limit {
		showbackRules = showbackRules[:config.Limit]
	}

	format.PrintResults(showbackRules,
		"id",
		"name",
		"metricName",
		"organizationName",
		"kind",
		"type",
		"globalAlertLimit",
		"projectAlertLimit",
		"price",
		"showbackCredentialName",
		"createdAt",
	)

	return
}
