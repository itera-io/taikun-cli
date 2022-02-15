package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"REMINDER", "reminder",
		),
		field.NewVisible(
			"SLACK-CONFIG", "slackConfigurationName",
		),
		field.NewHiddenWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
	},
)

type AddOptions struct {
	Emails               []string
	Name                 string
	OrganizationID       int32
	Reminder             string
	SlackConfigurationID int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if err := cmdutils.CheckFlagValue("reminder", opts.Reminder, types.AlertingReminders); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")
	cmd.Flags().Int32VarP(&opts.SlackConfigurationID, "slack-configuration-id", "s", 0, "Slack configuration ID")
	cmd.Flags().StringSliceVarP(&opts.Emails, "emails", "e", []string{}, "Emails")

	cmd.Flags().StringVarP(&opts.Reminder, "reminder", "r", "none", "Reminder")
	cmdutils.SetFlagCompletionValues(&cmd, "reminder", types.AlertingReminders.Keys()...)

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
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
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
