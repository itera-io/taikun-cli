package list

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/config"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
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
		Short: "List an alerting profile's webhooks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return cmderr.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return cmderr.OutputFormatInvalidError
			}
			alertingProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.AlertingProfileID = alertingProfileID
			return listRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")

	return cmd
}

func printResults(alertingWebhooks []*models.AlertingWebhookDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(alertingWebhooks)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(alertingWebhooks))
		for i, alertingWebhook := range alertingWebhooks {
			data[i] = alertingWebhook
		}
		format.PrettyPrintTable(data,
			"url",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_profiles.NewAlertingProfilesListParams().WithV(apiconfig.Version)
	params = params.WithID(&opts.AlertingProfileID)

	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesList(params, apiClient)
	if err != nil {
		return err
	}
	if len(response.Payload.Data) != 1 {
		return fmt.Errorf("Alerting profile with ID %d not found.", opts.AlertingProfileID)
	}
	alertingWebhooks := response.Payload.Data[0].Webhooks

	if opts.Limit != 0 && int32(len(alertingWebhooks)) > opts.Limit {
		alertingWebhooks = alertingWebhooks[:opts.Limit]
	}

	printResults(alertingWebhooks)
	return
}
