package showbackcredential

import (
	"taikun-cli/cmd/showbackcredential/create"
	"taikun-cli/cmd/showbackcredential/list"

	"github.com/spf13/cobra"
)

func NewCmdShowbackCredential() *cobra.Command {
	cmd := cobra.Command{
		Use:     "showback-credential <command>",
		Short:   "Manage showback credentials",
		Aliases: []string{"sbc"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
