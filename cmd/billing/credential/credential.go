package credential

import (
	"github.com/itera-io/taikun-cli/cmd/billing/credential/add"
	"github.com/itera-io/taikun-cli/cmd/billing/credential/list"
	"github.com/itera-io/taikun-cli/cmd/billing/credential/lock"
	"github.com/itera-io/taikun-cli/cmd/billing/credential/remove"
	"github.com/itera-io/taikun-cli/cmd/billing/credential/unlock"
	"github.com/spf13/cobra"
)

func NewCmdCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "credential <command>",
		Short:   "Manage billing credentials",
		Aliases: []string{"c"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
