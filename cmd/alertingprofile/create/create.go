package create

import (
	"taikun-cli/api"
	"taikun-cli/utils"

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
			if !utils.MapContains(utils.AlertingReminders, opts.Reminder) {
				return utils.UnknownFlagValueError(
					"reminder",
					opts.Reminder,
					utils.MapKeys(utils.AlertingReminders),
				)
			}
			return createRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().Int32VarP(&opts.SlackConfigurationID, "slack-configuration-id", "s", 0, "Slack configuration ID")
	cmd.Flags().StringSliceVarP(&opts.Emails, "emails", "e", []string{}, "Emails")
	cmd.Flags().StringVarP(&opts.Reminder, "reminder", "r", "None", "Reminder")
	utils.RegisterStaticFlagCompletion(cmd, "reminder", utils.MapKeys(utils.AlertingReminders)...)

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
		Reminder:             utils.GetAlertingReminder(opts.Reminder),
		SlackConfigurationID: opts.SlackConfigurationID,
	}

	if len(opts.Emails) != 0 {
		emails := make([]*models.AlertingEmailDto, len(opts.Emails))
		for i, email := range opts.Emails {
			emails[i].Email = email
		}
		body.Emails = emails
	}

	params := alerting_profiles.NewAlertingProfilesCreateParams().WithV(utils.ApiVersion).WithBody(&body)
	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesCreate(params, apiClient)
	if err == nil {
		utils.PrettyPrintJson(response.Payload)
	}

	return
}
