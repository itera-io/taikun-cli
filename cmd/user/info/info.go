package info

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

var infoFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"USERNAME", "username",
		),
		field.NewVisible(
			"EMAIL", "email",
		),
		field.NewVisible(
			"ROLE", "role",
		),
		field.NewHidden(
			"DISPLAY-NAME", "displayName",
		),
		field.NewHidden(
			"EMAIL-CONFIRMED", "isEmailConfirmed",
		),
		field.NewHidden(
			"EMAIL-NOTIFICATIONS", "isEmailNotificationEnabled",
		),
		field.NewHidden(
			"MUST-RESET-PASSWORD", "isForcedToResetPassword",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"OWNER", "owner",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED", "createdAt", out.FormatDateTimeString,
		),
	},
)

func NewCmdInfo() *cobra.Command {
	cmd := cobra.Command{
		Use:   "info",
		Short: "Retrieve information about the current user",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return infoRun(cmd)
		},
	}

	return &cmd
}

func infoRun(cmd *cobra.Command) error {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.UsersAPI.UsersUserInfo(ctx).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data.Data, infoFields)
}
