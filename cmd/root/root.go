package root

import (
	"fmt"
	"taikun-cli/cmd/accessprofile"
	"taikun-cli/cmd/backupcredential"
	"taikun-cli/cmd/billingcredential"
	"taikun-cli/cmd/cloudcredential"
	"taikun-cli/cmd/flavor"
	"taikun-cli/cmd/kubernetesprofile"
	"taikun-cli/cmd/organization"
	"taikun-cli/cmd/user"
	"taikun-cli/config"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "taikun <command> <subcommand> [flags]",
		Short:        "Taikun CLI",
		Long:         `Manage Taikun resources from the command line.`,
		SilenceUsage: true,
	}

	cmd.PersistentFlags().StringVarP(&config.OutputFormat, "format", "f", config.OutputFormatTable,
		fmt.Sprintf(
			"Output format: \"%s\" or \"%s\"",
			config.OutputFormatJson, config.OutputFormatTable,
		),
	)

	cmd.PersistentFlags().BoolVar(&config.ShowLargeValues, "show-large-values", false,
		"Prevent trimming of large cell values")

	cmd.AddCommand(accessprofile.NewCmdAccessProfile())
	cmd.AddCommand(backupcredential.NewCmdBackupCredential())
	cmd.AddCommand(billingcredential.NewCmdBillingCredential())
	cmd.AddCommand(cloudcredential.NewCmdBillingCredential())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(kubernetesprofile.NewCmdKubernetesProfile())
	cmd.AddCommand(organization.NewCmdOrganization())
	cmd.AddCommand(user.NewCmdUser())

	return cmd
}
