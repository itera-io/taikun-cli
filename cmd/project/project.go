package project

import (
	"taikun-cli/cmd/project/quotas"

	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}

	cmd.AddCommand(quotas.NewCmdQuotas())

	return cmd
}
