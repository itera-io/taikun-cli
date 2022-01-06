package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	AlertingProfileID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <alerting-profile-id>",
		Short: "List an alerting profile's integrations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alertingProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.AlertingProfileID = alertingProfileID
			return listRun(&opts)
		},
	}

	cmdutils.AddLimitFlag(cmd)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_integrations.NewAlertingIntegrationsListParams().WithV(apiconfig.Version)
	params = params.WithAlertingProfileID(opts.AlertingProfileID)

	response, err := apiClient.Client.AlertingIntegrations.AlertingIntegrationsList(params, apiClient)
	if err != nil {
		return err
	}
	alertingIntegrations := response.Payload

	if config.Limit != 0 && int32(len(alertingIntegrations)) > config.Limit {
		alertingIntegrations = alertingIntegrations[:config.Limit]
	}

	format.PrintResults(alertingIntegrations,
		"id",
		"alertingProfileName",
		"url",
		"token",
		"alertingIntegrationType",
	)
	return
}
