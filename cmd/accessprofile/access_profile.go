package accessprofile

import (
	"github.com/itera-io/taikun-cli/cmd/accessprofile/add"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/allowedhost"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/list"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/accessprofile/remove"
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

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(sshuser.NewCmdSshUser())
	cmd.AddCommand(allowedhost.NewCmdAllowedHost())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
