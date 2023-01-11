package usertoken

import (
	"github.com/itera-io/taikun-cli/cmd/usertoken/add"
	"github.com/spf13/cobra"
)

func NewCmdUserToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usertoken <command>",
		Short: "Manage user tokens",
	}

	cmd.AddCommand(add.NewCmdAdd())

	return cmd
}
