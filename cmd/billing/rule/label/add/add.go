package add

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	BillingRuleID int32
	Label         string
	Value         string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <billing-rule-id>",
		Short: "Add a label to a billing rule",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.BillingRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Label, "label", "l", "", "Label (required")
	cmdutils.MarkFlagRequired(&cmd, "label")

	cmd.Flags().StringVarP(&opts.Value, "value", "v", "", "Value (required")
	cmdutils.MarkFlagRequired(&cmd, "value")

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return err
	}

	body := models.RuleForUpdateDto{
		LabelsToAdd: []*models.PrometheusLabelListDto{
			{
				Label: opts.Label,
				Value: opts.Value,
			},
		},
	}

	params := prometheus.NewPrometheusUpdateParams().WithV(taikungoclient.Version)
	params = params.WithID(opts.BillingRuleID).WithBody(&body)

	_, err = apiClient.Client.Prometheus.PrometheusUpdate(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
