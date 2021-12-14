package kubernetesprofile

import (
	"taikun-cli/cmd/kubernetesprofile/create"
	"taikun-cli/cmd/kubernetesprofile/delete"
	"taikun-cli/cmd/kubernetesprofile/list"

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
	cmd.AddCommand(delete.NewCmdDelete())

	return cmd
}
