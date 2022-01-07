package list

import (
	"fmt"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	AlertingProfileID int32
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

	if config.Limit != 0 && int32(len(alertingWebhooks)) > config.Limit {
		alertingWebhooks = alertingWebhooks[:config.Limit]
	}

	format.PrintResults(alertingWebhooks,
		"url",
	)
	return
}
