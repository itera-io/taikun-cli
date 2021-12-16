package alertingprofile

import (
	"taikun-cli/cmd/alertingprofile/create"
	"taikun-cli/cmd/alertingprofile/delete"
	"taikun-cli/cmd/alertingprofile/list"

	"github.com/spf13/cobra"
)

func NewCmdAlertingProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerting-profile <command>",
		Short:   "Manage alerting profiles",
		Aliases: []string{"alert"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
