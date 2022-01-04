package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	DisplayName    string
	Email          string
	OrganizationID int32
	Role           string
	Username       string
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <username>",
		Short: "Create user",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Username = args[0]
			if !types.MapContains(types.UserRoles, opts.Role) {
				return types.UnknownFlagValueError(
					"role",
					opts.Role,
					types.MapKeys(types.UserRoles),
				)
			}
			return createRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email (required)")
	cmdutils.MarkFlagRequired(cmd, "email")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	cmdutils.MarkFlagRequired(cmd, "role")
	cmdutils.RegisterStaticFlagCompletion(cmd, "role", types.MapKeys(types.UserRoles)...)

	cmd.Flags().StringVarP(&opts.DisplayName, "display-name", "d", "", "Display name")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
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

	params := users.NewUsersCreateParams().WithV(apiconfig.Version).WithBody(body)
	response, err := apiClient.Client.Users.UsersCreate(params, apiClient)
	if err == nil {
		format.PrintResult(response.Payload,
			"id",
			"username",
			"role",
			"organizationName",
			"email",
			"isEmailConfirmed",
			"isEmailNotificationEnabled",
		)
	}

	return
}
