package showback

import (
	"github.com/itera-io/taikun-cli/cmd/showback/credential"
	"github.com/itera-io/taikun-cli/cmd/showback/rule"
	"github.com/spf13/cobra"
)

func NewCmdShowback() *cobra.Command {
	cmd := cobra.Command{
		Use:     "showback <command>",
		Short:   "Manage showback credentials and rules",
		Aliases: []string{"sb"},
	}

	cmd.AddCommand(credential.NewCmdCredential())
	cmd.AddCommand(rule.NewCmdRule())

	return &cmd
}
