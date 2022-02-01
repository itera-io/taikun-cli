package image

import (
	"github.com/itera-io/taikun-cli/cmd/project/image/bind"
	"github.com/itera-io/taikun-cli/cmd/project/image/list"
	"github.com/itera-io/taikun-cli/cmd/project/image/unbind"
	"github.com/spf13/cobra"
)

func NewCmdImage() *cobra.Command {
	cmd := cobra.Command{
		Use:   "image <command>",
		Short: "Manage a project's images",
	}

	cmd.AddCommand(bind.NewCmdBind())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(unbind.NewCmdUnbind())

	return &cmd
}
