package credential

import (
	"github.com/itera-io/taikun-cli/cmd/showback/credential/add"
	"github.com/itera-io/taikun-cli/cmd/showback/credential/delete"
	"github.com/itera-io/taikun-cli/cmd/showback/credential/list"
	"github.com/itera-io/taikun-cli/cmd/showback/credential/lock"
	"github.com/itera-io/taikun-cli/cmd/showback/credential/unlock"
	"github.com/spf13/cobra"
)

func NewCmdCredential() *cobra.Command {
	cmd := cobra.Command{
		Use:     "credential <command>",
		Short:   "Manage showback credentials",
		Aliases: []string{"c"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return &cmd
}
