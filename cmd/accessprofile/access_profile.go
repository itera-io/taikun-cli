package accessprofile

import (
	"github.com/itera-io/taikun-cli/cmd/accessprofile/create"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/delete"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/list"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/sshuser"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/unlock"

	"github.com/spf13/cobra"
)

func NewCmdAccessProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "access-profile <command>",
		Short:   "Manage access profiles",
		Aliases: []string{"acp"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(sshuser.NewCmdSshUser())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
