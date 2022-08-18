package allowedhost

import (
	"github.com/itera-io/taikun-cli/cmd/accessprofile/allowedhost/add"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/allowedhost/list"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/allowedhost/remove"
	"github.com/spf13/cobra"
)

func NewCmdAllowedHost() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "allowed-host <command>",
		Short:   "Manage an access profile's allowed hosts",
		Aliases: []string{"host"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return cmd
}
