package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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

	cmd := &cobra.Command{
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
	cmdutils.MarkFlagRequired(cmd, "email")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	cmdutils.MarkFlagRequired(cmd, "role")
	cmdutils.RegisterStaticFlagCompletion(cmd, "role", types.UserRoles.Keys()...)

	cmd.Flags().StringVarP(&opts.DisplayName, "display-name", "d", "", "Display name")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
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
		out.PrintResult(response.Payload,
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
