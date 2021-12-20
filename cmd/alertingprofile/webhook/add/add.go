package add

import (
	"errors"
	"fmt"
	"strings"
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	AlertingProfileID int32
	URL               string
	Headers           []string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <alerting-profile-id>",
		Short: "Add a webhook to an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alertingProfileID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			opts.AlertingProfileID = alertingProfileID
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.URL, "url", "u", "", "URL (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	cmd.Flags().StringSliceVarP(&opts.Headers, "headers", "H", []string{}, "Headers (format: \"key=value,key2=value2,...\")")

	return cmd
}

func getAlertingProfileWebhooks(id int32) ([]*models.AlertingWebhookDto, error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return nil, err
	}

	params := alerting_profiles.NewAlertingProfilesListParams().WithV(apiconfig.Version)
	params = params.WithID(&id)

	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesList(params, apiClient)
	if err != nil {
		return nil, err
	}
	if len(response.Payload.Data) != 1 {
		return nil, fmt.Errorf("Alerting profile with ID %d not found.", id)
	}
	return response.Payload.Data[0].Webhooks, nil
}

func parseAddOptions(opts *AddOptions) (*models.AlertingWebhookDto, error) {
	alertingWebhook := &models.AlertingWebhookDto{
		URL: opts.URL,
	}
	headers := make([]*models.WebhookHeaderDto, len(opts.Headers))
	for i, header := range opts.Headers {
		if len(header) == 0 {
			return nil, errors.New("Invalid empty webhook header")
		}
		tokens := strings.Split(header, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid webhook header format: %s", header)
		}
		headers[i] = &models.WebhookHeaderDto{
			Key:   tokens[0],
			Value: tokens[1],
		}
	}
	alertingWebhook.Headers = headers
	return alertingWebhook, nil
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	alertingWebhooks, err := getAlertingProfileWebhooks(opts.AlertingProfileID)
	if err != nil {
		return
	}

	newAlertingWebhook, err := parseAddOptions(opts)
	if err != nil {
		return
	}

	alertingWebhooks = append(alertingWebhooks, newAlertingWebhook)
	params := alerting_profiles.NewAlertingProfilesAssignWebhooksParams().WithV(apiconfig.Version)
	params = params.WithID(opts.AlertingProfileID)
	params = params.WithBody(alertingWebhooks)

	_, err = apiClient.Client.AlertingProfiles.AlertingProfilesAssignWebhooks(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
