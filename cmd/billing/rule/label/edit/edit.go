package edit

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	BillingRuleID int32
	Labels        []string
}

func NewCmdEdit() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "edit <billing-rule-id>",
		Short: "Edit labels of a billing rule",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.BillingRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringSliceVarP(&opts.Labels, "labels", "l", []string{}, "Labels (format: \"label=value,label2=value2,...\")")
	cmdutils.MarkFlagRequired(&cmd, "labels")

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// List other information that needs to be send
	data, response, err := myApiClient.Client.PrometheusRulesAPI.PrometheusrulesList(context.TODO()).Id(opts.BillingRuleID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Prepare the arguments for the query
	body := taikuncore.RuleForUpdateDto{
		Name:                  *taikuncore.NewNullableString(&data.GetData()[0].Name),
		MetricName:            data.GetData()[0].MetricName,
		Type:                  &data.GetData()[0].Type,
		Price:                 *taikuncore.NewNullableFloat64(&data.GetData()[0].Price),
		OperationCredentialId: data.GetData()[0].OperationCredential.OperationCredentialId,
	}

	body.Labels, err = out.ParseLabelsFlag(opts.Labels)
	if err != nil {
		return
	}

	// Execute a query into the API + graceful exit
	_, response, err = myApiClient.Client.PrometheusRulesAPI.PrometheusrulesUpdate(context.TODO(), opts.BillingRuleID).RuleForUpdateDto(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
