package delete

import (
	"errors"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/project/server/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
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
		Short: "Delete some or all servers from a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			if opts.DeleteAll {
				if len(opts.ServerIDs) != 0 {
					return errors.New("Cannot set both --server-ids and --all flags")
				}
			} else {
				if len(opts.ServerIDs) == 0 {
					return errors.New("Must set one of --server-ids and --all flags")
				}
			}
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32SliceVarP(&opts.ServerIDs, "server-ids", "s", []int32{}, "IDs of the servers to delete")
	cmd.Flags().BoolVarP(&opts.DeleteAll, "all", "a", false, "Delete all servers")

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteServerCommand{
		ProjectID: opts.ProjectID,
	}

	if opts.DeleteAll {
		allServers, err := list.ListServers(&list.ListOptions{ProjectID: opts.ProjectID})
		if err != nil {
			return err
		}
		allServerIDs := make([]int32, len(allServers))
		for i, server := range allServers {
			allServerIDs[i] = server.ID
		}
		body.ServerIds = allServerIDs
	} else {
		body.ServerIds = opts.ServerIDs
	}

	params := servers.NewServersDeleteParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, _, err = apiClient.Client.Servers.ServersDelete(params, apiClient)
	if err == nil {
		for _, id := range body.ServerIds {
			out.PrintDeleteSuccess("Server", id)
		}
	}

	return
}
