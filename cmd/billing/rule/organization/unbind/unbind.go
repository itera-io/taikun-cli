package unbind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UnbindOptions struct {
	BillingRuleID  int32
	OrganizationID int32
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := cobra.Command{
		Use:   "unbind <billing-rule-id>",
		Short: "Unbind a billing rule from an organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.BillingRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return unbindRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "organization-id")

	return &cmd
}

func unbindRun(opts *UnbindOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	isBound := false
	body := taikuncore.BindPrometheusOrganizationsCommand{
		PrometheusRuleId: &opts.BillingRuleID,
		Organizations: []taikuncore.BindOrganizationsToRuleDto{
			{
				OrganizationId: &opts.OrganizationID,
				IsBound:        &isBound,
			},
		},
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.PrometheusRulesAPI.PrometheusrulesBindOrganizations(context.TODO()).BindPrometheusOrganizationsCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.BindPrometheusOrganizationsCommand{
			PrometheusRuleID: opts.BillingRuleID,
			Organizations: []*models.BindOrganizationsToRuleDto{
				{
					OrganizationID: opts.OrganizationID,
					IsBound:        false,
				},
			},
		}

		params := prometheus.NewPrometheusBindOrganizationsParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.Prometheus.PrometheusBindOrganizations(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
