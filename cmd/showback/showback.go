package showback

import (
	"taikun-cli/cmd/showback/credential"

	"github.com/spf13/cobra"
)

func NewCmdShowback() *cobra.Command {
	cmd := cobra.Command{
		Use:     "showback <command>",
		Short:   "Manage showback credentials and rules",
		Aliases: []string{"sb"},
	}

	cmd.AddCommand(credential.NewCmdCredential())
	// TODO add Rule command

	return &cmd
}
