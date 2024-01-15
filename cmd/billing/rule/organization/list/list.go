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
			"ORG", "name",
		),
		field.NewVisible(
			"ORG-ID", "id",
		),
		field.NewVisible(
			"DISCOUNT-RATE", "ruleDiscountRate",
		),
		field.NewVisible(
			"GLOBAL-DISCOUNT-RATE", "globalDiscountRate",
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
		Short: "List a billing rule's bound organizations",
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

	boundOrganizations := data.Data[0].BoundOrganizations

	if opts.Limit != 0 && int32(len(boundOrganizations)) > opts.Limit {
		boundOrganizations = boundOrganizations[:opts.Limit]
	}

	return out.PrintResults(boundOrganizations, listFields)

}
