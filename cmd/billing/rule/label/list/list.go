package list

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
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
				return cmderr.ErrIDArgumentNotANumber
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.PrometheusRulesAPI.PrometheusrulesList(context.TODO()).Id(opts.BillingRuleID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	if len(data.Data) != 1 {
		return cmderr.ResourceNotFoundError("Billing rule", opts.BillingRuleID)
	}

	labels := data.Data[0].Labels

	if opts.Limit != 0 && int32(len(labels)) > opts.Limit {
		labels = labels[:opts.Limit]
	}

	return out.PrintResults(labels, listFields)
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := prometheus.NewPrometheusListOfRulesParams().WithV(taikungoclient.Version)
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
	*/
}
