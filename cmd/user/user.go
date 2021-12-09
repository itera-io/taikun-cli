package user

import (
	"taikun-cli/api"
	"taikun-cli/cmd/user/bind"

	"github.com/spf13/cobra"
)

func NewCmdUser(apiClient *api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user <command>",
		Short: "Manage users",
	}

	cmd.AddCommand(bind.NewCmdBind(apiClient))

	return cmd
}
