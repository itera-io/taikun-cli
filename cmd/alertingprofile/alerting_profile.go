package alertingprofile

import (
	"taikun-cli/cmd/alertingprofile/list"

	"github.com/spf13/cobra"
)

func NewCmdAlertingProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "alerting-profile",
		Short:   "Manage alerting profiles",
		Aliases: []string{"alert"},
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
