package create

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Emails               []string
	Name                 string
	OrganizationID       int32
	Reminder             string
	SlackConfigurationID int32
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if !types.MapContains(types.AlertingReminders, opts.Reminder) {
				return types.UnknownFlagValueError(
					"reminder",
					opts.Reminder,
					types.MapKeys(types.AlertingReminders),
				)
			}
			return createRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().Int32VarP(&opts.SlackConfigurationID, "slack-configuration-id", "s", 0, "Slack configuration ID")
	cmd.Flags().StringSliceVarP(&opts.Emails, "emails", "e", []string{}, "Emails")
	cmd.Flags().StringVarP(&opts.Reminder, "reminder", "r", "None", "Reminder")
	cmdutils.RegisterStaticFlagCompletion(cmd, "reminder", types.MapKeys(types.AlertingReminders)...)

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateAlertingProfileCommand{
		Name:                 opts.Name,
		OrganizationID:       opts.OrganizationID,
		Reminder:             types.GetAlertingReminder(opts.Reminder),
		SlackConfigurationID: opts.SlackConfigurationID,
	}

	if len(opts.Emails) != 0 {
		fmt.Println(opts.Emails)
		emails := make([]*models.AlertingEmailDto, len(opts.Emails))
		for i, email := range opts.Emails {
			emails[i] = &models.AlertingEmailDto{Email: email}
		}
		body.Emails = emails
	}

	params := alerting_profiles.NewAlertingProfilesCreateParams().WithV(utils.ApiVersion).WithBody(&body)
	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesCreate(params, apiClient)
	if err == nil {
		format.PrettyPrintJson(response.Payload)
	}

	return
}
