package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "username",
		),
		field.NewHidden(
			"DISPLAY-NAME", "displayName",
		),
		field.NewVisible(
			"ROLE", "role",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"OWNER", "owner",
		),
		field.NewVisible(
			"EMAIL", "email",
		),
		field.NewHidden(
			"EMAIL-CONFIRMED", "isEmailConfirmed",
		),
		field.NewVisible(
			"EMAIL-NOTIFICATIONS", "isEmailNotificationEnabled",
		),
		field.NewVisible(
			"APPROVED-BY-PARTNER", "isApprovedByPartner",
		),
		field.NewVisible(
			"CSM", "isCsm",
		),
		field.NewHidden(
			"SUBSCRIPTION-UPDATES", "isEligibleUpdateSubscription",
		),
		field.NewVisible(
			"MUST-RESET-PASSWORD", "isForcedToResetPassword",
		),
		field.NewHiddenWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"READ-ONLY", "isReadOnly",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED", "createdAt", out.FormatDateTimeString,
		),
	},
)

type AddOptions struct {
	DisplayName    string
	Email          string
	OrganizationID int32
	Role           string
	Username       string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <username>",
		Short: "Add a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Username = args[0]
			if err := cmdutils.CheckFlagValue("role", opts.Role, types.UserRoles); err != nil {
				return err
			}
			return addRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email (required)")
	cmdutils.MarkFlagRequired(&cmd, "email")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	cmdutils.MarkFlagRequired(&cmd, "role")
	cmdutils.SetFlagCompletionValues(&cmd, "role", types.UserRoles.Keys()...)

	cmd.Flags().StringVarP(&opts.DisplayName, "display-name", "d", "", "Display name")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CreateUserCommand{
		DisplayName:    opts.DisplayName,
		Email:          opts.Email,
		OrganizationID: opts.OrganizationID,
		Role:           types.GetUserRole(opts.Role),
		Username:       opts.Username,
	}

	params := users.NewUsersCreateParams().WithV(api.Version).WithBody(body)
	response, err := apiClient.Client.Users.UsersCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
