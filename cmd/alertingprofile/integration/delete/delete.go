package delete

import (
	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <alerting-integration-id>",
		Short: "Delete an alerting integration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := utils.Atoi32(args[0])
			if err != nil {
				return utils.WrongIDArgumentFormatError
			}
			return deleteRun(id)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_integrations.NewAlertingIntegrationsDeleteParams().WithV(utils.ApiVersion).WithID(id)
	_, _, err = apiClient.Client.AlertingIntegrations.AlertingIntegrationsDelete(params, apiClient)
	if err == nil {
		utils.PrintDeleteSuccess("Alerting integration", id)
	}

	return
}
