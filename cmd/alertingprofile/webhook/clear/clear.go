package clear

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdClear() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear <alerting-integration-id>",
		Short: "Clear an alerting profile's webhooks",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			return clearRun(id)
		},
	}

	return cmd
}

func clearRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	emptyWebhookList := make([]*models.AlertingWebhookDto, 0)
	params := alerting_profiles.NewAlertingProfilesAssignWebhooksParams().WithV(apiconfig.Version)
	params = params.WithID(id).WithBody(emptyWebhookList)

	_, err = apiClient.Client.AlertingProfiles.AlertingProfilesAssignWebhooks(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}