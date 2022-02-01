package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"LABEL", "label",
		),
		field.NewVisible(
			"VALUE", "value",
		),
	},
)

type ListOptions struct {
	BillingRuleID int32
	Limit         int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <billing-rule-id>",
		Short: "List a billing rule's labels",
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

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := prometheus.NewPrometheusListOfRulesParams().WithV(api.Version)
	params = params.WithID(&opts.BillingRuleID)

	response, err := apiClient.Client.Prometheus.PrometheusListOfRules(params, apiClient)
	if err != nil {
		return
	}
	if len(response.Payload.Data) != 1 {
		return cmderr.ResourceNotFoundError("Billing rule", opts.BillingRuleID)
	}

	labels := response.Payload.Data[0].Labels

	if opts.Limit != 0 && int32(len(labels)) > opts.Limit {
		labels = labels[:opts.Limit]
	}

	return out.PrintResults(labels, listFields)
}
