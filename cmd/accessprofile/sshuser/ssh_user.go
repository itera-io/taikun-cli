package sshuser

import (
	"taikun-cli/cmd/accessprofile/sshuser/create"
	"taikun-cli/cmd/accessprofile/sshuser/list"

	"github.com/spf13/cobra"
)

func NewCmdSshUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ssh-user <command>",
		Short:   "Manage SSH users",
		Aliases: []string{"ssh"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
