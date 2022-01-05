package billing

import (
	"github.com/itera-io/taikun-cli/cmd/billing/credential"
	"github.com/spf13/cobra"
)

func NewCmdBilling() *cobra.Command {
	cmd := cobra.Command{
		Use:   "billing <command>",
		Short: "Manage billing rules and credentials",
	}

	cmd.AddCommand(credential.NewCmdCredential())

	return &cmd
}
