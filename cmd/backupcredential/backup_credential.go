package backupcredential

import "github.com/spf13/cobra"

func NewCmdBackupCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup-credential <command>",
		Short:   "Manage backup credentials",
		Aliases: []string{"bc"},
	}

	// TODO add subcommands

	return cmd
}
