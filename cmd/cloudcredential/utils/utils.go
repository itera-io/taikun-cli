package utils

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	tk "github.com/itera-io/taikungoclient"
)

const (
	AWS = iota
	AZURE
	OPENSTACK
	GOOGLE
)

func GetCloudType(cloudCredentialID int32) (cloudType int, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.CloudCredentialAPI.CloudcredentialsDashboardList(context.TODO()).Id(cloudCredentialID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	} else {
		if len(data.GetAmazon()) == 1 {
			cloudType = AWS
		} else if len(data.GetAzure()) == 1 {
			cloudType = AZURE
		} else if len(data.GetOpenstack()) == 1 {
			cloudType = OPENSTACK
		} else if len(data.GetGoogle()) == 1 {
			cloudType = GOOGLE
		} else {
			err = cmderr.ResourceNotFoundError("Cloud credential", cloudCredentialID)
		}
	}
	return

}
