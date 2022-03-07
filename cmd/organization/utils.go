package organization

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikungoclient/client/users"
)

func GetDefaultOrganizationID() (id int32, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersDetailsParams().WithV(api.Version)

	response, err := apiClient.Client.Users.UsersDetails(params, apiClient)
	if err == nil {
		id = response.Payload.Data.OrganizationID
	}

	return
}
