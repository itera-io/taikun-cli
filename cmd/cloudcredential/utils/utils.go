package utils

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/cloud_credentials"
)

const (
	AWS = iota
	AZURE
	OPENSTACK
	GOOGLE
)

func GetCloudType(cloudCredentialID int32) (cloudType int, err error) {
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
}
