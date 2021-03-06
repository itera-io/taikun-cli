package remove

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <alerting-integration-id>...",
		Short: "Delete one or more alerting integrations",
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

	return cmd
}

func deleteRun(alertingIntegrationID int32) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := alerting_integrations.NewAlertingIntegrationsDeleteParams().WithV(taikungoclient.Version).WithID(alertingIntegrationID)

	_, _, err = apiClient.Client.AlertingIntegrations.AlertingIntegrationsDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Alerting integration", alertingIntegrationID)
	}

	return
}
