package complete

import (
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/openstack"
	"github.com/itera-io/taikungoclient/models"
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
}
