package list

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"URL", "url",
		),
	},
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
			alertingProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			opts.AlertingProfileID = alertingProfileID
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(cmd, &opts.Limit)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := alerting_profiles.NewAlertingProfilesListParams().WithV(taikungoclient.Version)
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

	return out.PrintResults(alertingWebhooks, listFields)
}
