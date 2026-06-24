package status

import (
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type StatusOptions struct {
	Mode string
}

func NewCmdRobotStatus() *cobra.Command {
	var opts StatusOptions

	cmd := cobra.Command{
		Use:   "status <ROBOT_ID>",
		Short: "Activate or deactivate robot user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return changeRobotStatus(cmd, args[0], &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Mode, "mode", "m", "", "Robot user mode")
	_ = cmd.MarkFlagRequired("mode")

	return &cmd
}

func changeRobotStatus(cmd *cobra.Command, robotID string, opts *StatusOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.RobotUserStatusManagementCommand{
		Id:   *taikuncore.NewNullableString(&robotID),
		Mode: *taikuncore.NewNullableString(&opts.Mode),
	}

	response, err := myApiClient.Client.RobotAPI.RobotStatus(ctx).RobotUserStatusManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
