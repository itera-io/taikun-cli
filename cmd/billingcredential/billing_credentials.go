package billingcredential

import (
	"taikun-cli/cmd/billingcredential/create"
	"taikun-cli/cmd/billingcredential/delete"
	"taikun-cli/cmd/billingcredential/list"

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
	cmd.AddCommand(delete.NewCmdDelete())

	return cmd
}
