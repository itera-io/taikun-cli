package cloudcredential

import (
	"taikun-cli/cmd/cloudcredential/delete"
	"taikun-cli/cmd/cloudcredential/lock"
	"taikun-cli/cmd/cloudcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdBillingCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-credential <command>",
		Short:   "Manage Cloud Credentials",
		Aliases: []string{"cloud"},
	}

	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
