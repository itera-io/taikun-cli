package credential

import (
	"taikun-cli/cmd/showback/credential/create"
	"taikun-cli/cmd/showback/credential/delete"
	"taikun-cli/cmd/showback/credential/list"
	"taikun-cli/cmd/showback/credential/lock"
	"taikun-cli/cmd/showback/credential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdCredential() *cobra.Command {
	cmd := cobra.Command{
		Use:     "credential <command>",
		Short:   "Manage showback credentials",
		Aliases: []string{"c", "cred"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return &cmd
}
