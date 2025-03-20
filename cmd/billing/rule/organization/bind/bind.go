package bind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	BillingRuleID  int32
	DiscountRate   float64
	OrganizationID int32
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := cobra.Command{
		Use:   "bind <billing-rule-id>",
		Short: "Bind a billing rule to an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.BillingRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return bindRun(&opts)
		},
	}

	cmd.Flags().Float64VarP(&opts.DiscountRate, "discount-rate", "d", 0, "Discount rate (required)")
	cmdutils.MarkFlagRequired(&cmd, "discount-rate")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "organization-id")

	return &cmd
}

func bindRun(opts *BindOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	billingRuleId := int32(opts.BillingRuleID)
	body := []taikuncore.AddOrganizationsToRuleDto{
		{
			OrganizationId:   &opts.OrganizationID,
			RuleDiscountRate: *taikuncore.NewNullableFloat64(&opts.DiscountRate),
		},
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.PrometheusRulesAPI.PrometheusrulesAddOrganizations(context.TODO(), billingRuleId).AddOrganizationsToRuleDto(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
