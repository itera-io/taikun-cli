package project

import (
	"taikun-cli/cmd/project/create"
	"taikun-cli/cmd/project/delete"
	"taikun-cli/cmd/project/list"
	"taikun-cli/cmd/project/lock"
	"taikun-cli/cmd/project/unlock"

	"github.com/spf13/cobra"
)

func NewCmdProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project <command>",
		Short: "Manage projects",
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
