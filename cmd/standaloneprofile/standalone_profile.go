package standaloneprofile

import (
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/add"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/delete"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/list"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/rename"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/securitygroup"
	"github.com/itera-io/taikun-cli/cmd/standaloneprofile/unlock"
	"github.com/spf13/cobra"
)

func NewCmdStandaloneProfile() *cobra.Command {
	cmd := cobra.Command{
		Use:     "standalone-profile <command>",
		Short:   "Manage standalone profiles",
		Aliases: []string{"sp"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(rename.NewCmdRename())
	cmd.AddCommand(securitygroup.NewCmdSecurityGroup())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return &cmd
}
