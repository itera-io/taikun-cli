package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"

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
			if err := cmdutils.CheckFlagValue("reminder", opts.Reminder, types.AlertingReminders); err != nil {
				return err
			}
			return createRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().Int32VarP(&opts.SlackConfigurationID, "slack-configuration-id", "s", 0, "Slack configuration ID")
	cmd.Flags().StringSliceVarP(&opts.Emails, "emails", "e", []string{}, "Emails")
	cmd.Flags().StringVarP(&opts.Reminder, "reminder", "r", "none", "Reminder")
	cmdutils.RegisterStaticFlagCompletion(cmd, "reminder", types.AlertingReminders.Keys()...)

	cmdutils.AddOutputOnlyIDFlag(cmd)

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
		emails := make([]*models.AlertingEmailDto, len(opts.Emails))
		for i, email := range opts.Emails {
			emails[i] = &models.AlertingEmailDto{Email: email}
		}
		body.Emails = emails
	}

	params := alerting_profiles.NewAlertingProfilesCreateParams().WithV(api.Version).WithBody(&body)
	response, err := apiClient.Client.AlertingProfiles.AlertingProfilesCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"organizationName",
			"slackConfigurationName",
			"reminder",
			"isLocked",
		)
	}

	return
}
