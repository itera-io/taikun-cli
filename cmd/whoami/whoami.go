package whoami

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/spf13/cobra"
)

func NewCmdWhoAmI() *cobra.Command {
	cmd := cobra.Command{
		Use:   "whoami",
		Short: "Print username",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return whoAmIRun()
		},
	}

	return &cmd
}

func whoAmIRun() (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDetailsParams().WithV(apiconfig.Version)

	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		format.Println(response.Payload.Data.Username)
	}

	return
}
