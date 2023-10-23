package root

import (
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/accessprofile"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile"
	"github.com/itera-io/taikun-cli/cmd/backupcredential"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization"
	"github.com/itera-io/taikun-cli/cmd/project"
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile"
	"github.com/itera-io/taikun-cli/cmd/user"
	"github.com/itera-io/taikun-cli/cmd/usertoken"
	"github.com/itera-io/taikun-cli/cmd/whoami"
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
				return cmderr.ErrUnknownOutputFormat
			}
			return nil
		},
	}

	setPersistentFlags(cmd)

	cmd.AddCommand(accessprofile.NewCmdAccessProfile())
	cmd.AddCommand(alertingprofile.NewCmdAlertingProfile())
	cmd.AddCommand(backupcredential.NewCmdBackupCredential())
	//cmd.AddCommand(billing.NewCmdBilling())
	cmd.AddCommand(cloudcredential.NewCmdCloudCredential())
	//cmd.AddCommand(kubernetesprofile.NewCmdKubernetesProfile())
	cmd.AddCommand(organization.NewCmdOrganization())
	//cmd.AddCommand(policyprofile.NewCmdPolicyProfile())
	cmd.AddCommand(project.NewCmdProject())
	//cmd.AddCommand(showback.NewCmdShowback())
	cmd.AddCommand(slackconfiguration.NewCmdSlackConfiguration())
	cmd.AddCommand(standaloneprofile.NewCmdStandaloneProfile())
	cmd.AddCommand(user.NewCmdUser())
	cmd.AddCommand(usertoken.NewCmdUserToken())
	//cmd.AddCommand(usertoken2.NewCmdUserToken2())
	cmd.AddCommand(whoami.NewCmdWhoami())

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
