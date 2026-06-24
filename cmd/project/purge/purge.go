package purge

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdPurge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "purge <project-id>",
		Short: "Deletes all k8s servers and all virtual clusters for this project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return purgeRun(cmd, args[0])
		},
		Args: cobra.ExactArgs(1),
	}

	return cmd
}

func purgeRun(cmd *cobra.Command, projectIdString string) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	// Project ID
	projectId, err := types.Atoi32(projectIdString)
	if err != nil {
		return err
	}

	// Server IDs
	data, response, err := myApiClient.Client.ServersAPI.ServersDetails(ctx, projectId).Execute()
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

	response, err = myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDelete(ctx).ProjectDeploymentDeleteServersCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Project", projectId)
	return nil
}
