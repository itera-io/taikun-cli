package showbackcredential

import (
	"taikun-cli/cmd/showbackcredential/create"
	"taikun-cli/cmd/showbackcredential/delete"
	"taikun-cli/cmd/showbackcredential/list"
	"taikun-cli/cmd/showbackcredential/lock"
	"taikun-cli/cmd/showbackcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdShowbackCredential() *cobra.Command {
	cmd := cobra.Command{
		Use:     "showback-credential <command>",
		Short:   "Manage showback credentials",
		Aliases: []string{"sbc"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return &cmd
}
