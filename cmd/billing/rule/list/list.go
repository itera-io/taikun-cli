package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdList() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "List billing rules",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun()
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddSortByAndReverseFlags(&cmd, models.AccessProfilesListDto{})

	return &cmd
}

func listRun() (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := prometheus.NewPrometheusListOfRulesParams().WithV(apiconfig.Version)
	if config.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var billingRules = make([]*models.PrometheusRuleListDto, 0)
	for {
		response, err := apiClient.Client.Prometheus.PrometheusListOfRules(params, apiClient)
		if err != nil {
			return err
		}
		billingRules = append(billingRules, response.Payload.Data...)
		count := int32(len(billingRules))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(billingRules)) > config.Limit {
		billingRules = billingRules[:config.Limit]
	}

	out.PrintResults(billingRules,
		"id",
		"name",
		"metricName",
		"price",
		"createdAt",
		"type",
	)

	return
}
