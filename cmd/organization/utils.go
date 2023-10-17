package organization

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
)

func GetDefaultOrganizationID() (id int32, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.UsersAPI.UsersUserInfo(context.TODO()).Execute()
	if err != nil {
		return -1, tk.CreateError(response, err)
	}
	if err == nil {
		id = data.Data.GetOrganizationId()
	}

	return
}
