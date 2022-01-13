package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
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
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddSortByAndReverseFlags(&cmd, models.AccessProfilesListDto{})

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := showback.NewShowbackRulesListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.ReverseSortDirection {
		api.ReverseSortDirection()
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(&api.SortDirection)
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

	out.PrintResults(showbackRules,
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
