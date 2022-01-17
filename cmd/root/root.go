package root

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/accessprofile"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"

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
			if config.Limit < 0 {
				return cmderr.NegativeLimitFlagError
			}
			return nil
		},
	}

	setPersistentFlags(cmd)

	cmd.AddCommand(accessprofile.NewCmdAccessProfile())
	// cmd.AddCommand(alertingprofile.NewCmdAlertingProfile()) TODO
	// cmd.AddCommand(backupcredential.NewCmdBackupCredential()) TODO
	// cmd.AddCommand(billing.NewCmdBilling()) TODO
	// cmd.AddCommand(cloudcredential.NewCmdCloudCredential()) TODO
	// cmd.AddCommand(kubernetesprofile.NewCmdKubernetesProfile()) TODO
	// cmd.AddCommand(organization.NewCmdOrganization()) TODO
	// cmd.AddCommand(policyprofile.NewCmdPolicyProfile()) TODO
	// cmd.AddCommand(project.NewCmdProject()) TODO
	// cmd.AddCommand(showback.NewCmdShowback()) TODO
	// cmd.AddCommand(slackconfiguration.NewCmdSlackConfiguration()) TODO
	// cmd.AddCommand(user.NewCmdUser()) TODO
	// cmd.AddCommand(whoami.NewCmdWhoAmI()) TODO

	return cmd
}

func setPersistentFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVar(
		&config.NoDecorate,
		"no-decorate",
		false,
		"Display output table without field names and separators",
	)

	cmd.PersistentFlags().StringVarP(
		&config.OutputFormat,
		"format", "F",
		config.OutputFormatTable,
		fmt.Sprintf("Output format for list-type commands: one of %v", config.OutputFormats),
	)
	cmdutils.SetFlagCompletionValues(cmd, "format", config.OutputFormats...)

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
