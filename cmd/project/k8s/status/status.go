package status

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/spf13/cobra"
)

type StatusOptions struct {
	ServerID int32
}

func NewCmdStatus() *cobra.Command {
	var opts StatusOptions

	cmd := cobra.Command{
		Use:   "status <server-id>",
		Short: "Display a Kubernetes server's status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ServerID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return statusRun(&opts)
		},
	}

	return &cmd
}

func statusRun(opts *StatusOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := servers.NewServersShowServerStatusParams().WithV(api.Version)
	params = params.WithServerID(opts.ServerID)

	response, err := apiClient.Client.Servers.ServersShowServerStatus(params, apiClient)
	if err == nil {
		out.Println(response.Payload)
	}

	return
}
