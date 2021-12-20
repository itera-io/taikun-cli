package showbackcredential

import "github.com/spf13/cobra"

func NewCmdShowbackCredential() *cobra.Command {
	cmd := cobra.Command{
		Use:     "showback-credential <command>",
		Short:   "Manage showback credentials",
		Aliases: []string{"sbc"},
	}

	// TODO add subcommands

	return &cmd
}
