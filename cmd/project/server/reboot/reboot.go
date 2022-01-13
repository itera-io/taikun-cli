package reboot

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/servers"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type RebootOptions struct {
	ServerID int32
}

func NewCmdReboot() *cobra.Command {
	var opts RebootOptions

	cmd := cobra.Command{
		Use:   "reboot <server-id>",
		Short: "Reboot a server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ServerID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return rebootRun(&opts)
		},
	}

	return &cmd
}

func rebootRun(opts *RebootOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.RebootServerCommand{
		ServerID: opts.ServerID,
	}

	params := servers.NewServersRebootParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Servers.ServersReboot(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}
	return
}
