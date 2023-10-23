package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultipleChildResources(opts.BillingRuleID, ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32VarP(&opts.BillingRuleID, "billing-rule-id", "b", 0, "Billing rule ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "billing-rule-id")

	return &cmd
}

func deleteRun(billingRuleID int32, labelID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.RuleForUpdateDto{
		LabelsToDelete: []taikuncore.PrometheusLabelDeleteDto{
			{
				Id: &billingRuleID,
			},
		},
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.PrometheusRulesAPI.PrometheusrulesUpdate(context.TODO(), billingRuleID).RuleForUpdateDto(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Billing rule label", labelID)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.RuleForUpdateDto{
			LabelsToDelete: []*models.PrometheusLabelDeleteDto{
				{
					ID: labelID,
				},
			},
		}

		params := prometheus.NewPrometheusUpdateParams().WithV(taikungoclient.Version)
		params.WithID(billingRuleID).WithBody(&body)

		_, err = apiClient.Client.Prometheus.PrometheusUpdate(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Billing rule label", labelID)
		}

		return
	*/
}
