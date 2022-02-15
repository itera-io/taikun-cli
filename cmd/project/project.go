package project

import (
	"github.com/itera-io/taikun-cli/cmd/project/add"
	"github.com/itera-io/taikun-cli/cmd/project/alert"
	"github.com/itera-io/taikun-cli/cmd/project/delete"
	"github.com/itera-io/taikun-cli/cmd/project/etc"
	"github.com/itera-io/taikun-cli/cmd/project/flavor"
	"github.com/itera-io/taikun-cli/cmd/project/image"
	"github.com/itera-io/taikun-cli/cmd/project/info"
	"github.com/itera-io/taikun-cli/cmd/project/k8s"
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig"
	"github.com/itera-io/taikun-cli/cmd/project/list"
	"github.com/itera-io/taikun-cli/cmd/project/lock"
	"github.com/itera-io/taikun-cli/cmd/project/monitoringtoggle"
	"github.com/itera-io/taikun-cli/cmd/project/quota"
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
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(etc.NewCmdEtc())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(image.NewCmdImage())
	cmd.AddCommand(info.NewCmdInfo())
	cmd.AddCommand(k8s.NewCmdK8s())
	cmd.AddCommand(kubeconfig.NewCmdKubeconfig())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(quota.NewCmdQuota())
	cmd.AddCommand(monitoringtoggle.NewCmdMonitoringToggle())
	cmd.AddCommand(unlock.NewCmdUnlock())
	cmd.AddCommand(vm.NewCmdVm())

	return cmd
}
