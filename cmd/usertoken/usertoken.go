package usertoken

import (
	"github.com/itera-io/taikun-cli/cmd/usertoken/add"
	"github.com/itera-io/taikun-cli/cmd/usertoken/bearer"
	"github.com/itera-io/taikun-cli/cmd/usertoken/bind"
	"github.com/itera-io/taikun-cli/cmd/usertoken/list"
	"github.com/itera-io/taikun-cli/cmd/usertoken/remove"
	"github.com/itera-io/taikun-cli/cmd/usertoken/showendpoints"
	"github.com/itera-io/taikun-cli/cmd/usertoken/unbind"
	"github.com/spf13/cobra"
)

func NewCmdUserToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "usertoken <command>",
		Short: "Manage user tokens",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(unbind.NewCmdUnbind())
	cmd.AddCommand(showendpoints.NewCmdShowendpoints())
	cmd.AddCommand(bearer.NewCmdBearer())

	return cmd
}
