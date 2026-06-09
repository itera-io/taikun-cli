package delete

import (
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdDeleteRobot() *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete <GROUP_ID>",
		Short: "Delete group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteRobot(cmd, args[0])
		},
	}

	return &cmd
}

func deleteRobot(cmd *cobra.Command, robotID string) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.RobotAPI.RobotDelete(ctx, robotID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("robot", robotID)
	return nil
}
