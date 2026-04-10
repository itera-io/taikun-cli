package delete

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDeleteRobot() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <GROUP_ID>",
		Short: "Delete group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteRobot(args[0])
		},
	}

	return &cmd
}

func deleteRobot(robotID string) (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.RobotAPI.RobotDelete(context.TODO(), robotID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("robot", robotID)
	return nil
}
