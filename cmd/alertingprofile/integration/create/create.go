package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	AlertingProfileID int32
	URL               string
	Type              string
	Token             string
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <alerting-profile-id>",
		Short: "Create an alerting integration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !types.MapContains(types.AlertingIntegrationTypes, opts.Type) {
				return types.UnknownFlagValueError(
					"type",
					opts.Type,
					types.MapKeys(types.AlertingIntegrationTypes),
				)
			}
			alertingProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.AlertingProfileID = alertingProfileID
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "URL (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	cmd.Flags().StringVarP(&opts.Type, "type", "t", "", "Type (required)")
	cmdutils.MarkFlagRequired(cmd, "type")
	cmdutils.RegisterStaticFlagCompletion(cmd, "type", types.MapKeys(types.AlertingIntegrationTypes)...)

	cmd.Flags().StringVar(&opts.Token, "token", "", "Token")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateAlertingIntegrationCommand{
		AlertingProfileID: opts.AlertingProfileID,
		AlertingIntegration: &models.AlertingIntegrationDto{
			URL:                     opts.URL,
			Token:                   opts.Token,
			AlertingIntegrationType: types.GetAlertingIntegrationType(opts.Type),
		},
	}

	params := alerting_integrations.NewAlertingIntegrationsCreateParams().WithV(apiconfig.Version).WithBody(&body)
	if response, err := apiClient.Client.AlertingIntegrations.AlertingIntegrationsCreate(params, apiClient); err == nil {
		format.PrintResult(response.Payload,
			"id",
			"alertingProfileName",
			"url",
			"token",
			"alertingIntegrationType",
		)
	}

	return
}
