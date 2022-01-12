package delete

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

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
				return cmderr.IDArgumentNotANumberError
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_integrations.NewAlertingIntegrationsDeleteParams().WithV(apiconfig.Version).WithID(id)
	_, _, err = apiClient.Client.AlertingIntegrations.AlertingIntegrationsDelete(params, apiClient)
	if err == nil {
		out.PrintDeleteSuccess("Alerting integration", id)
	}

	return
}
