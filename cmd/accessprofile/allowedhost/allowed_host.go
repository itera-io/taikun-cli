package allowedhost

import (
	"github.com/itera-io/taikun-cli/cmd/accessprofile/allowedhost/add"
	"github.com/spf13/cobra"
)

func NewCmdSshUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "allowed-host <command>",
		Short:   "Manage an access profile's allowed hosts",
		Aliases: []string{"host"},
	}

	cmd.AddCommand(add.NewCmdAdd())

	return cmd
}
