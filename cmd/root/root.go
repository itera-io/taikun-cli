package root

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/accessprofile"
	"github.com/itera-io/taikun-cli/cmd/alertingprofile"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/cmdutils/options"
	"github.com/itera-io/taikun-cli/cmd/organization"
	"github.com/itera-io/taikun-cli/cmd/whoami"

	"github.com/spf13/cobra"
)

type RootOptions struct {
	OutputFormat string
	Quiet        bool
}

func (opts *RootOptions) GetOutputFormatOption() *string {
	return &opts.OutputFormat
}

func (opts *RootOptions) GetQuietOption() *bool {
	return &opts.Quiet
}

func NewCmdRoot() *cobra.Command {
	var opts RootOptions

	cmd := &cobra.Command{
		Use:          "taikun <command> <subcommand> [flags]",
		Short:        "Taikun CLI",
		Long:         `Manage Taikun resources from the command line.`,
		SilenceUsage: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if !options.OutputFormatIsValid() {
				return cmderr.OutputFormatInvalidError
			}
			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(
		&opts.OutputFormat,
		"format", "F",
		options.OutputFormatTable,
		fmt.Sprintf("Output format for list-type commands: one of %v",
			[]string{
				options.OutputFormatJson,
				options.OutputFormatTable,
			},
		),
	)
	cmdutils.SetFlagCompletionValues(cmd, "format",
		options.OutputFormatJson,
		options.OutputFormatTable,
	)

	cmd.PersistentFlags().BoolVarP(
		opts.Quiet,
		"quiet", "q",
		false,
		"Turn off output (does not disable output to stderr)",
	)

	cmd.AddCommand(accessprofile.NewCmdAccessProfile(&opts))
	cmd.AddCommand(alertingprofile.NewCmdAlertingProfile(&opts))
	// cmd.AddCommand(backupcredential.NewCmdBackupCredential()) TODO
	// cmd.AddCommand(billing.NewCmdBilling()) TODO
	// cmd.AddCommand(cloudcredential.NewCmdCloudCredential()) TODO
	// cmd.AddCommand(kubernetesprofile.NewCmdKubernetesProfile()) TODO
	cmd.AddCommand(organization.NewCmdOrganization(&opts))
	// cmd.AddCommand(policyprofile.NewCmdPolicyProfile()) TODO
	// cmd.AddCommand(project.NewCmdProject()) TODO
	// cmd.AddCommand(showback.NewCmdShowback()) TODO
	// cmd.AddCommand(slackconfiguration.NewCmdSlackConfiguration()) TODO
	// cmd.AddCommand(user.NewCmdUser()) TODO
	cmd.AddCommand(whoami.NewCmdWhoAmI(&opts))

	return cmd
}
