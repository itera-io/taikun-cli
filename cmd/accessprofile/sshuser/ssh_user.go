package sshuser

import (
	"github.com/itera-io/taikun-cli/cmd/accessprofile/sshuser/add"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/sshuser/list"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/sshuser/remove"
	"github.com/spf13/cobra"
)

func NewCmdSshUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ssh-user <command>",
		Short:   "Manage an access profile's SSH users",
		Aliases: []string{"ssh"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return cmd
}
