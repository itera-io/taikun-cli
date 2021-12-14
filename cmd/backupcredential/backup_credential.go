package backupcredential

import (
	"taikun-cli/cmd/backupcredential/create"
	"taikun-cli/cmd/backupcredential/list"

	"github.com/spf13/cobra"
)

func NewCmdBackupCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup-credential <command>",
		Short:   "Manage backup credentials",
		Aliases: []string{"bc"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
