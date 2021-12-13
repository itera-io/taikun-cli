package accessprofile

import (
	"taikun-cli/cmd/accessprofile/create"
	"taikun-cli/cmd/accessprofile/delete"
	"taikun-cli/cmd/accessprofile/list"
	"taikun-cli/cmd/accessprofile/sshuser"

	"github.com/spf13/cobra"
)

func NewCmdAccessProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "access-profile <command>",
		Short:   "Manage access profiles",
		Aliases: []string{"access"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(sshuser.NewCmdSshUser())

	return cmd
}
