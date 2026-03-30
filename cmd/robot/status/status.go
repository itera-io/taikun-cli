package status

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
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
			return changeRobotStatus(args[0], &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Mode, "mode", "m", "", "Robot user mode")
	_ = cmd.MarkFlagRequired("mode")

	return &cmd
}

func changeRobotStatus(robotID string, opts *StatusOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.RobotUserStatusManagementCommand{
		Id:   *taikuncore.NewNullableString(&robotID),
		Mode: *taikuncore.NewNullableString(&opts.Mode),
	}

	response, err := myApiClient.Client.RobotAPI.RobotStatus(context.TODO()).RobotUserStatusManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
