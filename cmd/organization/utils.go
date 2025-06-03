package organization

import (
	"context"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
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

func GetOrganizationIDFromCloudCredential(ccid int32, client *tk.Client) (int32, error) {
	data, response, err := client.Client.CloudCredentialAPI.CloudcredentialsOrgList(context.TODO()).IsAdmin(false).Id(ccid).Execute()
	if err != nil {
		return -1, tk.CreateError(response, err)
	} else if len(data) != 1 {
		return -1, fmt.Errorf("invalid cloud credential id %d", ccid)
	}
	return data[0].GetOrganizationId(), nil
}
