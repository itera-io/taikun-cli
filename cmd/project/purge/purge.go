package purge

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func NewCmdPurge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "purge <project-id>",
		Short: "Deletes all k8s servers and all virtual clusters for this project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return purgeRun(args[0])
		},
		Args: cobra.ExactArgs(1),
	}

	return cmd
}

func purgeRun(projectIdString string) (err error) {
	myApiClient := tk.NewClient()

	// Project ID
	projectId, err := types.Atoi32(projectIdString)
	if err != nil {
		return err
	}

	// Server IDs
	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(context.TODO(), projectId).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	var serversToDelete []int32
	for _, server := range data.GetData() {
		serversToDelete = append(serversToDelete, server.Id)
	}

	alwaysTellThe := true // https://www.youtube.com/watch?v=GvlN1Lr3yt8
	body := taikuncore.ProjectDeploymentDeleteServersCommand{
		ProjectId:                &projectId,
		ServerIds:                serversToDelete,
		ForceDeleteVClusters:     &alwaysTellThe,
		DeleteAutoscalingServers: &alwaysTellThe,
	}

	_, response, err = myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDelete(context.TODO()).ProjectDeploymentDeleteServersCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Project", projectId)
	return nil
}
