package alert

import (
	"github.com/itera-io/taikun-cli/cmd/project/alert/attach"
	"github.com/itera-io/taikun-cli/cmd/project/alert/detach"
	"github.com/spf13/cobra"
)

func NewCmdAlert() *cobra.Command {
	cmd := cobra.Command{
		Use:   "alert <command>",
		Short: "Manage a project's alerting profile",
	}

	cmd.AddCommand(attach.NewCmdAttach())
	cmd.AddCommand(detach.NewCmdDetach())

	return &cmd
}
