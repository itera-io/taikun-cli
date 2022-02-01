package slackconfiguration

import (
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration/add"
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration/delete"
	"github.com/itera-io/taikun-cli/cmd/slackconfiguration/list"
	"github.com/spf13/cobra"
)

func NewCmdSlackConfiguration() *cobra.Command {
	cmd := cobra.Command{
		Use:     "slack-configuration <command>",
		Short:   "Manage Slack configurations",
		Aliases: []string{"slack"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return &cmd
}
