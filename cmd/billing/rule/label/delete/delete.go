package delete

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	BillingRuleID int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <label-id>...",
		Short: "Delete one or more labels from a billing rule",
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return cmdutils.DeleteMultipleChildResources(opts.BillingRuleID, ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32VarP(&opts.BillingRuleID, "billing-rule-id", "b", 0, "Billing rule ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "billing-rule-id")

	return &cmd
}

func deleteRun(billingRuleID int32, id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.RuleForUpdateDto{
		LabelsToDelete: []*models.PrometheusLabelDeleteDto{
			{
				ID: id,
			},
		},
	}

	params := prometheus.NewPrometheusUpdateParams().WithV(apiconfig.Version)
	params.WithID(billingRuleID).WithBody(&body)

	_, err = apiClient.Client.Prometheus.PrometheusUpdate(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Billing rule label", id)
	}

	return
}
