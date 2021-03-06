package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
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
				return cmderr.ErrIDArgumentNotANumber
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

func getAlertingProfileWebhooks(alertingProfileID int32) ([]*models.AlertingWebhookDto, error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return nil, err
	}

	params := alerting_profiles.NewAlertingProfilesListParams().WithV(taikungoclient.Version)
	params = params.WithID(&alertingProfileID)

	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesList(params, apiClient)
	if err != nil {
		return nil, err
	}

	if len(response.Payload.Data) != 1 {
		return nil, fmt.Errorf("Alerting profile with ID %d not found.", alertingProfileID)
	}

	return response.Payload.Data[0].Webhooks, nil
}

func parseAddOptions(opts *AddOptions) (*models.AlertingWebhookDto, error) {
	alertingWebhook := &models.AlertingWebhookDto{
		URL: opts.URL,
	}

	headers := make([]*models.WebhookHeaderDto, len(opts.Headers))

	for headerIndex, header := range opts.Headers {
		if len(header) == 0 {
			return nil, errors.New("Invalid empty webhook header")
		}

		tokens := strings.Split(header, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid webhook header format: %s", header)
		}

		headers[headerIndex] = &models.WebhookHeaderDto{
			Key:   tokens[0],
			Value: tokens[1],
		}
	}

	alertingWebhook.Headers = headers

	return alertingWebhook, nil
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
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
	params := alerting_profiles.NewAlertingProfilesAssignWebhooksParams().WithV(taikungoclient.Version)
	params = params.WithID(opts.AlertingProfileID)
	params = params.WithBody(alertingWebhooks)

	_, err = apiClient.Client.AlertingProfiles.AlertingProfilesAssignWebhooks(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
