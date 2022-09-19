package kubeconfig

import (
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig/add"
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig/download"
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig/list"
	"github.com/itera-io/taikun-cli/cmd/project/kubeconfig/remove"
	"github.com/spf13/cobra"
)

func NewCmdKubeconfig() *cobra.Command {
	cmd := cobra.Command{
		Use:   "kubeconfig <command>",
		Short: "Manage kubeconfigs",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(download.NewCmdDownload())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return &cmd
}
