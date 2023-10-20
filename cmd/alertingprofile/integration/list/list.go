package list

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"ALERTING-PROFILE", "alertingProfileName",
		),
		field.NewVisible(
			"URL", "url",
		),
		field.NewVisible(
			"TOKEN", "token",
		),
		field.NewVisible(
			"TYPE", "alertingIntegrationType",
		),
	},
)

type ListOptions struct {
	AlertingProfileID int32
	Limit             int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <alerting-profile-id>",
		Short: "List an alerting profile's integrations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alertingProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.AlertingProfileID = alertingProfileID
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	alertingIntegrations, response, err := myApiClient.Client.AlertingIntegrationsAPI.AlertingintegrationsList(context.TODO(), opts.AlertingProfileID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	if opts.Limit != 0 && int32(len(alertingIntegrations)) > opts.Limit {
		alertingIntegrations = alertingIntegrations[:opts.Limit]
	}

	return out.PrintResults(alertingIntegrations, listFields)
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := alerting_integrations.NewAlertingIntegrationsListParams().WithV(taikungoclient.Version)
		params = params.WithAlertingProfileID(opts.AlertingProfileID)

		response, err := apiClient.Client.AlertingIntegrations.AlertingIntegrationsList(params, apiClient)
		if err != nil {
			return err
		}

		alertingIntegrations := response.Payload
		if opts.Limit != 0 && int32(len(alertingIntegrations)) > opts.Limit {
			alertingIntegrations = alertingIntegrations[:opts.Limit]
		}

		return out.PrintResults(alertingIntegrations, listFields)
	*/
}
