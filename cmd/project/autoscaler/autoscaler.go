package autoscaler

import (
	"github.com/itera-io/taikun-cli/cmd/project/autoscaler/disable"
	"github.com/spf13/cobra"
)

func NewCmdAutoscaler() *cobra.Command {
	cmd := cobra.Command{
		Use:   "autoscaler <command>",
		Short: "Manage a project's autoscaling",
	}

	cmd.AddCommand(disable.NewCmdDisable())

	return &cmd
}
