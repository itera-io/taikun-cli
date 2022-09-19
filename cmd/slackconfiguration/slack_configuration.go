package slackconfiguration

import (
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration/add"
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration/list"
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration/remove"
	"github.com/spf13/cobra"
)

func NewCmdSlackConfiguration() *cobra.Command {
	cmd := cobra.Command{
		Use:     "slack-configuration <command>",
		Short:   "Manage Slack configurations",
		Aliases: []string{"slack"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(remove.NewCmdDelete())

	return &cmd
}
