package trustedregistry

import (
	"github.com/itera-io/taikun-cli/cmd/accessprofile/trustedregistry/add"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/trustedregistry/list"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/trustedregistry/remove"
	"github.com/spf13/cobra"
)

func NewCmdTrustedRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "trusted-registry <command>",
		Short:   "Manage an access profile's trusted registries",
		Aliases: []string{"host"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return cmd
}
