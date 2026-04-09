package add

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

// addFields defines a slice of fields corresponding to .
// Some columns are set as visible by default and some are hidden by default.
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
	DisplayName string
	Email       string
	AccountID   int32
	Username    string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <username>",
		Short: "Add a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Username = args[0]
			return addRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	// Email is a required flag
	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email (required)")
	cmdutils.MarkFlagRequired(&cmd, "email")

	// Display name optional flag. Default none.
	cmd.Flags().StringVarP(&opts.DisplayName, "display-name", "d", "", "Display name")

	// Account ID mandatory flag. Default 0.
	cmd.Flags().Int32VarP(&opts.AccountID, "account-id", "", -1, "Account ID")
	cmdutils.MarkFlagRequired(&cmd, "account-id")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

// addRun calls the API with a custom body from arguments. It than prints the result.
func addRun(opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.CreateUserCommand{}
	body.SetDisplayName(opts.DisplayName)
	body.SetEmail(opts.Email)
	body.SetAccountId(opts.AccountID)
	body.SetUsername(opts.Username)

	data, response, err := myApiClient.Client.UsersAPI.UsersCreate(context.TODO()).CreateUserCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, addFields)
}
