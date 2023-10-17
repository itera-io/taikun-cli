package status

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
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
	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ServersAPI.ServersStatus(context.TODO(), opts.ServerID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.Println(data)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := servers.NewServersShowServerStatusParams().WithV(taikungoclient.Version)
		params = params.WithServerID(opts.ServerID)

		response, err := apiClient.Client.Servers.ServersShowServerStatus(params, apiClient)
		if err == nil {
			out.Println(response.Payload)
		}

		return*/
}
