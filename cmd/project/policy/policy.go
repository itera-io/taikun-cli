package policy

import (
	"github.com/itera-io/taikun-cli/cmd/project/policy/disable"
	"github.com/itera-io/taikun-cli/cmd/project/policy/enable"
	"github.com/spf13/cobra"
)

func NewCmdPolicy() *cobra.Command {
	cmd := cobra.Command{
		Use:   "policy <command>",
		Short: "Manage a project's policy profile",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enable.NewCmdEnable())

	return &cmd
}
