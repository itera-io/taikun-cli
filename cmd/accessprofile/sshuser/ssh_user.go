package sshuser

import "github.com/spf13/cobra"

func NewCmdSshUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ssh-user <command>",
		Short:   "Manage SSH users",
		Aliases: []string{"ssh"},
	}

	// TODO add subcommands

	return cmd
}
