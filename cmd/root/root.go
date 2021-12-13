package root

import (
	"taikun-cli/cmd/accessprofile"
	"taikun-cli/cmd/billingcredential"
	"taikun-cli/cmd/flavor"
	"taikun-cli/cmd/noop"
	"taikun-cli/cmd/organization"
	"taikun-cli/cmd/user"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "taikun <command> <subcommand> [flags]",
		Short:        "Taikun CLI",
		Long:         `Manage Taikun resources from the command line.`,
		SilenceUsage: true,
	}

	cmd.AddCommand(accessprofile.NewCmdAccessProfile())
	cmd.AddCommand(billingcredential.NewCmdBillingCredential())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(noop.NewCmdNoop())
	cmd.AddCommand(organization.NewCmdOrganization())
	cmd.AddCommand(user.NewCmdUser())

	return cmd
}
