package project

import (
	"taikun-cli/cmd/project/list"

	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
