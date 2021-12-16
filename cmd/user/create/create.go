package create

import (
	"taikun-cli/api"
	"taikun-cli/utils"

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
			return createRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringVarP(&opts.Email, "email", "e", "", "Email (required)")
	utils.MarkFlagRequired(cmd, "email")

	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role (required)")
	utils.MarkFlagRequired(cmd, "role")

	cmd.Flags().StringVarP(&opts.DisplayName, "display-name", "d", "", "Display name")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

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
		Role:           utils.GetUserRole(opts.Role),
		Username:       opts.Username,
	}

	params := users.NewUsersCreateParams().WithV(utils.ApiVersion).WithBody(body)
	response, err := apiClient.Client.Users.UsersCreate(params, apiClient)
	if err == nil {
		utils.PrettyPrintJson(response.Payload)
	}

	return
}
