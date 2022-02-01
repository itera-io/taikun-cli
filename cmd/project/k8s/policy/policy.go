package policy

import (
	"github.com/itera-io/taikun-cli/cmd/project/k8s/policy/disable"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/policy/enable"
	"github.com/spf13/cobra"
)

func NewCmdPolicy() *cobra.Command {
	cmd := cobra.Command{
		Use:   "policy <command>",
		Short: "Manage Kubernetes servers' policy",
	}

	cmd.AddCommand(disable.NewCmdDisable())
	cmd.AddCommand(enable.NewCmdEnable())

	return &cmd
}
