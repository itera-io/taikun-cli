package add

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"strings"
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

func getAlertingProfileWebhooks(alertingProfileID int32) ([]taikuncore.AlertingWebhookDto, error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesList(context.TODO()).Id(alertingProfileID).Execute()
	if err != nil {
		return nil, tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	if len(data.Data) != 1 {
		return nil, fmt.Errorf("Alerting profile with ID %d not found.", alertingProfileID)
	}

	return data.Data[0].Webhooks, nil

}

func parseAddOptions(opts *AddOptions) (*taikuncore.AlertingWebhookDto, error) {
	alertingWebhook := &taikuncore.AlertingWebhookDto{
		Url: *taikuncore.NewNullableString(&opts.URL),
	}

	headers := make([]taikuncore.WebhookHeaderDto, len(opts.Headers))

	for headerIndex, header := range opts.Headers {
		if len(header) == 0 {
			return nil, fmt.Errorf("Invalid empty webhook header")
		}

		tokens := strings.Split(header, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid webhook header format: %s", header)
		}

		headers[headerIndex] = taikuncore.WebhookHeaderDto{
			Key:   *taikuncore.NewNullableString(&tokens[0]),
			Value: *taikuncore.NewNullableString(&tokens[1]),
		}
	}
	alertingWebhook.Headers = headers

	return alertingWebhook, nil

}

func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()

	alertingWebhooks, err := getAlertingProfileWebhooks(opts.AlertingProfileID)
	if err != nil {
		return
	}

	newAlertingWebhook, err := parseAddOptions(opts)
	if err != nil {
		return
	}

	alertingWebhooks = append(alertingWebhooks, *newAlertingWebhook)

	_, response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesAssignWebhooks(context.TODO(), opts.AlertingProfileID).AlertingWebhookDto(alertingWebhooks).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
