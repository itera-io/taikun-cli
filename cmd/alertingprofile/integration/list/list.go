package list

import (
	"taikun-cli/api"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	AlertingProfileID int32
	Limit             int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <alerting-profile-id>",
		Short: "List an alerting profile's integrations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return utils.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			alertingProfileID, err := utils.Atoi32(args[0])
			if err != nil {
				return utils.WrongIDArgumentFormatError
			}
			opts.AlertingProfileID = alertingProfileID
			return listRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")

	return cmd
}

func printResults(alertingIntegrations []*models.AlertingIntegrationsListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		utils.PrettyPrintJson(alertingIntegrations)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(alertingIntegrations))
		for i, alertingIntegration := range alertingIntegrations {
			data[i] = alertingIntegration
		}
		utils.PrettyPrintTable(data,
			"id",
			"alertingProfileName",
			"url",
			"token",
			"alertingIntegrationType",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_integrations.NewAlertingIntegrationsListParams().WithV(utils.ApiVersion)
	params = params.WithAlertingProfileID(opts.AlertingProfileID)

	response, err := apiClient.Client.AlertingIntegrations.AlertingIntegrationsList(params, apiClient)
	if err != nil {
		return err
	}
	alertingIntegrations := response.Payload

	if opts.Limit != 0 && int32(len(alertingIntegrations)) > opts.Limit {
		alertingIntegrations = alertingIntegrations[:opts.Limit]
	}

	printResults(alertingIntegrations)
	return
}
