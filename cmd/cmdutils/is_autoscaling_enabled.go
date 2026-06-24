package cmdutils

import (
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func IsAutoscalingEnabled(cmd *cobra.Command, projectID int32) (autoscalingEnabled bool, err error) {
	ctx, cancel := APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(ctx, projectID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	autoscalingEnabled = data.Project.GetIsAutoscalingEnabled()
	return
}
