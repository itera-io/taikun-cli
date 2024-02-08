package project

import (
	"github.com/itera-io/taikun-cli/cmd/project/add"
	"github.com/itera-io/taikun-cli/cmd/project/alert"
	"github.com/itera-io/taikun-cli/cmd/project/autoscaler"
	"github.com/itera-io/taikun-cli/cmd/project/backup"
	"github.com/itera-io/taikun-cli/cmd/project/backupsource"
	"github.com/itera-io/taikun-cli/cmd/project/disablemonitoring"
	"github.com/itera-io/taikun-cli/cmd/project/enablemonitoring"
	"github.com/itera-io/taikun-cli/cmd/project/etc"
	"github.com/itera-io/taikun-cli/cmd/project/flavor"
	"github.com/itera-io/taikun-cli/cmd/project/image"
	"github.com/itera-io/taikun-cli/cmd/project/info"
	"github.com/itera-io/taikun-cli/cmd/project/k8s"
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig"
	"github.com/itera-io/taikun-cli/cmd/project/list"
	"github.com/itera-io/taikun-cli/cmd/project/lock"
	"github.com/itera-io/taikun-cli/cmd/project/quota"
	"github.com/itera-io/taikun-cli/cmd/project/remove"
	"github.com/itera-io/taikun-cli/cmd/project/restore"
	"github.com/itera-io/taikun-cli/cmd/project/set"
	"github.com/itera-io/taikun-cli/cmd/project/spot"
	"github.com/itera-io/taikun-cli/cmd/project/unlock"
	"github.com/itera-io/taikun-cli/cmd/project/vm"
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
	cmd.AddCommand(etc.NewCmdEtc())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(image.NewCmdImage())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(k8s.NewCmdK8s())
	cmd.AddCommand(kubeconfig.NewCmdKubeconfig())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(enablemonitoring.NewCmdEnableMonitoring())
	cmd.AddCommand(disablemonitoring.NewCmdDisableMonitoring())
	cmd.AddCommand(quota.NewCmdQuota())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())
	cmd.AddCommand(vm.NewCmdVm())
	cmd.AddCommand(restore.NewCmdRestore())
	cmd.AddCommand(backupsource.NewCmdBackupSource())
	cmd.AddCommand(autoscaler.NewCmdAutoscaler())
	cmd.AddCommand(set.NewCmdSet())
	cmd.AddCommand(spot.NewCmdSpot())

	return cmd
}
