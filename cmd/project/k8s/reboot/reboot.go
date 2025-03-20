package reboot

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type RebootOptions struct {
	ServerID   int32
	RebootType bool
}

func NewCmdReboot() *cobra.Command {
	var opts RebootOptions

	cmd := cobra.Command{
		Use:   "reboot <server-id>",
		Short: "Reboot a Kubernetes server",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ServerID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return rebootRun(&opts)
		},
	}
	cmd.Flags().BoolVarP(&opts.RebootType, "hard", "f", false, "Force hard reboot of server")

	return &cmd
}

func rebootRun(opts *RebootOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.RebootServerCommand{}
	body.SetServerId(opts.ServerID)
	if opts.RebootType {
		body.SetType("hard")
	} else {
		body.SetType("soft")
	}

	_, response, err := myApiClient.Client.ServersAPI.ServersReboot(context.TODO()).RebootServerCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
