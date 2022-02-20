package policyprofile

import (
	"github.com/itera-io/taikun-cli/cmd/policyprofile/add"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/list"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/remove"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/unlock"
	"github.com/spf13/cobra"
)

func NewCmdPolicyProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "policy-profile <command>",
		Short:   "Manage policy profiles",
		Aliases: []string{"pp"},
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
