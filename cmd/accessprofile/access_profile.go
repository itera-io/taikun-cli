package accessprofile

import (
	"taikun-cli/cmd/accessprofile/create"
	"taikun-cli/cmd/accessprofile/delete"
	"taikun-cli/cmd/accessprofile/list"

	"github.com/spf13/cobra"
)

func NewCmdAccessProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "access-profile <command>",
		Short: "Manage access profiles",
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(delete.NewCmdDelete())

	return cmd
}
