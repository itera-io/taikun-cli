package kubernetesprofile

import (
	"github.com/itera-io/taikun-cli/cmd/kubernetesprofile/create"
	"github.com/itera-io/taikun-cli/cmd/kubernetesprofile/delete"
	"github.com/itera-io/taikun-cli/cmd/kubernetesprofile/list"
	"github.com/itera-io/taikun-cli/cmd/kubernetesprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/kubernetesprofile/unlock"

	"github.com/spf13/cobra"
)

func NewCmdKubernetesProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "kubernetes-profile <command>",
		Short:   "Manage kubernetes profiles",
		Aliases: []string{"kp"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
