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

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := cloud_credentials.NewCloudCredentialsDashboardListParams().WithV(taikungoclient.Version)
		params = params.WithID(&cloudCredentialID)

		response, err := apiClient.Client.CloudCredentials.CloudCredentialsDashboardList(params, apiClient)
		if err == nil {
			if len(response.Payload.Amazon) == 1 {
				cloudType = AWS
			} else if len(response.Payload.Azure) == 1 {
				cloudType = AZURE
			} else if len(response.Payload.Openstack) == 1 {
				cloudType = OPENSTACK
			} else if len(response.Payload.Google) == 1 {
				cloudType = GOOGLE
			} else {
				err = cmderr.ResourceNotFoundError("Cloud credential", cloudCredentialID)
			}
		}

		return
	*/
}
