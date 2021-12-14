package backupcredential

import (
	"taikun-cli/cmd/backupcredential/create"
	"taikun-cli/cmd/backupcredential/delete"
	"taikun-cli/cmd/backupcredential/list"
	"taikun-cli/cmd/backupcredential/lock"
	"taikun-cli/cmd/backupcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdBackupCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup-credential <command>",
		Short:   "Manage backup credentials",
		Aliases: []string{"bc"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
