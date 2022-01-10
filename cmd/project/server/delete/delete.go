package delete

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	ProjectID int32
	ServerIDs []int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := cobra.Command{
		Use:   "delete <server-id>...",
		Short: "Delete one or more servers from a project",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ServerIDs, err = cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return deleteRun(&opts)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "project-id")

	return &cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteServerCommand{
		ProjectID: opts.ProjectID,
		ServerIds: opts.ServerIDs,
	}

	params := servers.NewServersDeleteParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	_, _, err = apiClient.Client.Servers.ServersDelete(params, apiClient)
	if err == nil {
		for _, id := range opts.ServerIDs {
			format.PrintDeleteSuccess("Server", id)
		}
	}

	return
}
