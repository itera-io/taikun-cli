package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <billing-rule-id>...",
		Short: "Delete one or more billing rules",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return &cmd
}

func deleteRun(billingRuleID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := prometheus.NewPrometheusDeleteParams().WithV(taikungoclient.Version)
	params = params.WithID(billingRuleID)

	_, err = apiClient.Client.Prometheus.PrometheusDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Billing rule", billingRuleID)
	}

	return
}
