package status

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return statusRun(cmd, &opts)
		},
	}

	return &cmd
}

func statusRun(cmd *cobra.Command, opts *StatusOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	data, response, err := myApiClient.Client.ServersAPI.ServersStatus(ctx, opts.ServerID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.Println(data)
	return

}
