package billingcredential

import (
	"taikun-cli/cmd/billingcredential/create"
	"taikun-cli/cmd/billingcredential/delete"
	"taikun-cli/cmd/billingcredential/list"
	"taikun-cli/cmd/billingcredential/lock"
	"taikun-cli/cmd/billingcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdBillingCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "billing-credential <command>",
		Short:   "Manage Billing Credentials",
		Aliases: []string{"operation-credential"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
