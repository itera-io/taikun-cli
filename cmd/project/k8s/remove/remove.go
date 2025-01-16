package remove

import (
	"context"
	"errors"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	DeleteAll bool
	ProjectID int32
	ServerIDs []int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete some or all Kubernetes servers from a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if opts.DeleteAll {
				if len(opts.ServerIDs) != 0 {
					return errors.New("Cannot set both --server-ids and --all-servers flags")
				}
			} else {
				if len(opts.ServerIDs) == 0 {
					return errors.New("Must set one of --server-ids and --all-servers flags")
				}
			}
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32SliceVarP(&opts.ServerIDs, "server-ids", "s", []int32{}, "IDs of the servers to delete")
	cmd.Flags().BoolVarP(&opts.DeleteAll, "all-servers", "a", false, "Delete all of the project's servers")

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.ProjectDeploymentDeleteServersCommand{
		ProjectId: &opts.ProjectID,
	}

	if opts.DeleteAll {
		allServers, err := list.ListServers(&list.ListOptions{ProjectID: opts.ProjectID})
		if err != nil {
			return err
		}

		if len(allServers) == 0 {
			return fmt.Errorf("project %d has no Kubernetes servers", opts.ProjectID)
		}

		allServerIDs := make([]int32, len(allServers))
		for i, server := range allServers {
			allServerIDs[i] = server.GetId()
		}
		body.SetServerIds(allServerIDs)
	} else {
		body.SetServerIds(opts.ServerIDs)
	}
	_, response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDelete(context.TODO()).ProjectDeploymentDeleteServersCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	for _, id := range body.ServerIds {
		out.PrintDeleteSuccess("Server", id)
	}
	return

}
