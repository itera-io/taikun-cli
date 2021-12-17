package kubernetesprofile

import (
	"taikun-cli/cmd/kubernetesprofile/create"
	"taikun-cli/cmd/kubernetesprofile/delete"
	"taikun-cli/cmd/kubernetesprofile/list"
	"taikun-cli/cmd/kubernetesprofile/lock"
	"taikun-cli/cmd/kubernetesprofile/unlock"

	"github.com/spf13/cobra"
)

func NewCmdKubernetesProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kubernetes-profile <command>",
		Short:   "Manage kubernetes profiles",
		Aliases: []string{"k8s-profile"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
