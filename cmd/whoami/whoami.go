package whoami

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient"
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
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDetailsParams().WithV(taikungoclient.Version)

	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		out.Println(response.Payload.Data.Username)
	}

	return
}
