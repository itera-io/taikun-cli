package k8s

import (
	"github.com/itera-io/taikun-cli/cmd/project/k8s/add"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/backup"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/commit"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/delete"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/list"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/policy"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/reboot"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/repair"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/status"
	"github.com/itera-io/taikun-cli/cmd/project/k8s/upgrade"
	"github.com/spf13/cobra"
)

func NewCmdK8s() *cobra.Command {
	cmd := cobra.Command{
		Use:     "k8s <command>",
		Short:   "Manage a project's Kubernetes servers",
		Aliases: []string{"kubernetes"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(backup.NewCmdBackup())
	cmd.AddCommand(commit.NewCmdCommit())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(policy.NewCmdPolicy())
	cmd.AddCommand(reboot.NewCmdReboot())
	cmd.AddCommand(repair.NewCmdRepair())
	cmd.AddCommand(status.NewCmdStatus())
	cmd.AddCommand(upgrade.NewCmdUpgrade())

	return &cmd
}
