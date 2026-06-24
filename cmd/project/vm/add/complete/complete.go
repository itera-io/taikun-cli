package complete

import (
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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

	volumeTypes, err := getOpenStackVolumeTypes(cmd, projectID)
	if err == nil {
		completions = append(completions, volumeTypes...)
	}

	return
}

func getOpenStackVolumeTypes(cmd *cobra.Command, projectID int32) (volumeTypes []string, err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.OpenstackVolumeTypeListQuery{
		ProjectId: *taikuncore.NewNullableInt32(&projectID),
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.OpenstackCloudCredentialAPI.OpenstackVolumes(ctx).OpenstackVolumeTypeListQuery(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	// Manipulate the gathered data
	volumeTypes = data
	return volumeTypes, nil

}
