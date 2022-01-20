package add

import (
	"errors"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_integrations"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"URL", "url",
		),
	},
)

type AddOptions struct {
	AlertingProfileID int32
	URL               string
	Type              string
	Token             string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <alerting-profile-id>",
		Short: "Add an integration to an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.AlertingProfileID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			if err := cmdutils.CheckFlagValue("type", opts.Type, types.AlertingIntegrationTypes); err != nil {
				return err
			}
			if opts.Type != types.AlertingIntegrationTypeTeams && opts.Token == "" {
				return errors.New("--token must be set with this integration type")
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "URL (required)")
	cmdutils.MarkFlagRequired(&cmd, "url")

	cmd.Flags().StringVarP(&opts.Type, "type", "t", "", "Type (required)")
	cmdutils.MarkFlagRequired(&cmd, "type")
	cmdutils.SetFlagCompletionValues(&cmd, "type", types.AlertingIntegrationTypes.Keys()...)

	cmd.Flags().StringVar(&opts.Token, "token", "", "Token")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
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

	params := alerting_integrations.NewAlertingIntegrationsCreateParams().WithV(api.Version).WithBody(&body)
	if response, err := apiClient.Client.AlertingIntegrations.AlertingIntegrationsCreate(params, apiClient); err == nil {
		out.PrintResult(response.Payload, addFields)
	}

	return
}
