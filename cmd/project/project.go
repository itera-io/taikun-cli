package project

import (
	"github.com/itera-io/taikun-cli/cmd/project/add"
	"github.com/itera-io/taikun-cli/cmd/project/alert"
	"github.com/itera-io/taikun-cli/cmd/project/backup"
	"github.com/itera-io/taikun-cli/cmd/project/commit"
	"github.com/itera-io/taikun-cli/cmd/project/delete"
	"github.com/itera-io/taikun-cli/cmd/project/etc"
	"github.com/itera-io/taikun-cli/cmd/project/flavor"
	"github.com/itera-io/taikun-cli/cmd/project/info"
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig"
	"github.com/itera-io/taikun-cli/cmd/project/list"
	"github.com/itera-io/taikun-cli/cmd/project/lock"
	"github.com/itera-io/taikun-cli/cmd/project/policy"
	"github.com/itera-io/taikun-cli/cmd/project/quotas"
	"github.com/itera-io/taikun-cli/cmd/project/repair"
	"github.com/itera-io/taikun-cli/cmd/project/server"
	"github.com/itera-io/taikun-cli/cmd/project/unlock"
	"github.com/itera-io/taikun-cli/cmd/project/upgrade"

	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(alert.NewCmdAlert())
	cmd.AddCommand(backup.NewCmdBackup())
	cmd.AddCommand(commit.NewCmdCommit())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(etc.NewCmdEtc())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(kubeconfig.NewCmdKubeconfig())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(policy.NewCmdPolicy())
	cmd.AddCommand(quotas.NewCmdQuotas())
	cmd.AddCommand(repair.NewCmdRepair())
	cmd.AddCommand(server.NewCmdServer())
	cmd.AddCommand(unlock.NewCmdUnlock())
	cmd.AddCommand(upgrade.NewCmdUpgrade())

	return cmd
}
