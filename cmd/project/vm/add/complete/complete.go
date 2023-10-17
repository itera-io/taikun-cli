package complete

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

func VolumeTypeCompletionFunc(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
	completions = make([]string, 0)

	if len(args) == 0 {
		return
	}

	projectID, err := types.Atoi32(args[0])
	if err != nil {
		return
	}

	volumeTypes, err := getOpenStackVolumeTypes(projectID)
	if err == nil {
		completions = append(completions, volumeTypes...)
	}

	return
}

func getOpenStackVolumeTypes(projectID int32) (volumeTypes []string, err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.OpenstackVolumeTypeListQuery{
		ProjectId: *taikuncore.NewNullableInt32(&projectID),
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.OpenstackCloudCredentialAPI.OpenstackVolumes(context.TODO()).OpenstackVolumeTypeListQuery(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	// Manipulate the gathered data
	volumeTypes = data
	return volumeTypes, nil
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.OpenstackVolumeTypeListQuery{ProjectID: projectID}
		params := openstack.NewOpenstackVolumeTypesParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		response, err := apiClient.Client.Openstack.OpenstackVolumeTypes(params, apiClient)
		if err == nil {
			volumeTypes = response.Payload
		}

		return
	*/
}
