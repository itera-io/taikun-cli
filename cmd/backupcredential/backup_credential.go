package backupcredential

import (
	"github.com/itera-io/taikun-cli/cmd/backupcredential/add"
	"github.com/itera-io/taikun-cli/cmd/backupcredential/delete"
	"github.com/itera-io/taikun-cli/cmd/backupcredential/list"
	"github.com/itera-io/taikun-cli/cmd/backupcredential/lock"
	"github.com/itera-io/taikun-cli/cmd/backupcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdBackupCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup-credential <command>",
		Short:   "Manage backup credentials",
		Aliases: []string{"bc"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
