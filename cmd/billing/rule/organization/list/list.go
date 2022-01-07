package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	BillingRuleID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <billing-rule-id>",
		Short: "List a billing rule's bound organizations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.BillingRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := prometheus.NewPrometheusListOfRulesParams().WithV(apiconfig.Version)
	params = params.WithID(&opts.BillingRuleID)

	response, err := apiClient.Client.Prometheus.PrometheusListOfRules(params, apiClient)
	if err != nil {
		return
	}
	if len(response.Payload.Data) != 1 {
		return cmderr.ResourceNotFoundError("Billing rule", opts.BillingRuleID)
	}

	boundOrganizations := response.Payload.Data[0].BoundOrganizations

	if config.Limit != 0 && int32(len(boundOrganizations)) > config.Limit {
		boundOrganizations = boundOrganizations[:config.Limit]
	}

	format.PrintResults(boundOrganizations,
		"organizationId",
		"organizationName",
		"ruleDiscountRate",
		"globalDiscountRate",
	)

	return
}
