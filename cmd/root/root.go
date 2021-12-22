package root

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/accessprofile"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile"
	"github.com/itera-io/taikun-cli/cmd/backupcredential"
	"github.com/itera-io/taikun-cli/cmd/billingcredential"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/flavor"
	"github.com/itera-io/taikun-cli/cmd/kubernetesprofile"
	"github.com/itera-io/taikun-cli/cmd/organization"
	"github.com/itera-io/taikun-cli/cmd/policyprofile"
	"github.com/itera-io/taikun-cli/cmd/project"
	"github.com/itera-io/taikun-cli/cmd/showback"
	"github.com/itera-io/taikun-cli/cmd/user"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/list"

	"github.com/spf13/cobra"
)

func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "taikun <command> <subcommand> [flags]",
		Short:        "Taikun CLI",
		Long:         `Manage Taikun resources from the command line.`,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if !config.OutputFormatIsValid() {
				return cmderr.OutputFormatInvalidError
			}
			if list.Limit < 0 {
				return cmderr.NegativeLimitFlagError
			}
			return nil
		},
	}

	setPersistentFlags(cmd)

	cmd.AddCommand(accessprofile.NewCmdAccessProfile())
	cmd.AddCommand(alertingprofile.NewCmdAlertingProfile())
	cmd.AddCommand(backupcredential.NewCmdBackupCredential())
	cmd.AddCommand(billingcredential.NewCmdBillingCredential())
	cmd.AddCommand(cloudcredential.NewCmdBillingCredential())
	cmd.AddCommand(flavor.NewCmdFlavor())
	cmd.AddCommand(kubernetesprofile.NewCmdKubernetesProfile())
	cmd.AddCommand(organization.NewCmdOrganization())
	cmd.AddCommand(policyprofile.NewCmdPolicyProfile())
	cmd.AddCommand(project.NewCmdProject())
	cmd.AddCommand(showback.NewCmdShowback())
	cmd.AddCommand(user.NewCmdUser())

	return cmd
}

func setPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(
		&config.OutputFormat,
		"format", "F",
		config.OutputFormatTable,
		fmt.Sprintf("Output format for list-type commands: one of %v", config.OutputFormats),
	)
	cmdutils.RegisterStaticFlagCompletion(cmd, "format", config.OutputFormats...)

	cmd.PersistentFlags().BoolVar(
		&config.ShowLargeValues,
		"show-large-values",
		false,
		"Prevent trimming of large cell values",
	)

	cmd.PersistentFlags().BoolVarP(
		&config.Quiet,
		"quiet", "q",
		false,
		"Turn off output (does not disable output to stderr)",
	)
}
