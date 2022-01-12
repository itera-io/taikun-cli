package policy

import (
	"github.com/itera-io/taikun-cli/cmd/project/policy/disable"
	"github.com/itera-io/taikun-cli/cmd/project/policy/enforce"
	"github.com/spf13/cobra"
)

func NewCmdPolicy() *cobra.Command {
	cmd := cobra.Command{
		Use:   "policy <command>",
		Short: "Manage a project's policy profile",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enforce.NewCmdEnforce())

	return &cmd
}
