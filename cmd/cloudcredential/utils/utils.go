package utils

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
)

func GetCloudType(cloudCredentialID int32) (cloudType taikuncore.CloudType, err error) {
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.CloudCredentialAPI.CloudcredentialsOrgList(context.TODO()).IsAdmin(false).Id(cloudCredentialID).Execute()

	if len(data) == 0 {
		return taikuncore.CLOUDTYPE_NONE, cmderr.ResourceNotFoundError("Cloud credential", cloudCredentialID)
	}

	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	switch data[0].GetCloudType() {
	case taikuncore.CLOUDTYPE_AWS:
		return taikuncore.CLOUDTYPE_AWS, nil
	case taikuncore.CLOUDTYPE_AZURE:
		return taikuncore.CLOUDTYPE_AZURE, nil
	case taikuncore.CLOUDTYPE_OPENSTACK:
		return taikuncore.CLOUDTYPE_OPENSTACK, nil
	case taikuncore.CLOUDTYPE_GOOGLE:
		return taikuncore.CLOUDTYPE_GOOGLE, nil
	case taikuncore.CLOUDTYPE_PROXMOX:
		return taikuncore.CLOUDTYPE_PROXMOX, nil
	case taikuncore.CLOUDTYPE_VSPHERE:
		return taikuncore.CLOUDTYPE_VSPHERE, nil
	case taikuncore.CLOUDTYPE_OPENSHIFT:
		return taikuncore.CLOUDTYPE_OPENSHIFT, nil
	default:
		return taikuncore.CLOUDTYPE_NONE, cmderr.ResourceNotFoundError("Cloud credential", cloudCredentialID)
	}

}
