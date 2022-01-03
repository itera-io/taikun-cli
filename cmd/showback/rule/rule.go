package rule

import "github.com/spf13/cobra"

func NewCmdRule() *cobra.Command {
	cmd := cobra.Command{
		Use:     "rule <command>",
		Short:   "Manage showback rules",
		Aliases: []string{"r"},
	}

	// TODO add subcommands

	return &cmd
}
