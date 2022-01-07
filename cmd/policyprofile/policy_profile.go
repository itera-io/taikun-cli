package policyprofile

import (
	"github.com/itera-io/taikun-cli/cmd/policyprofile/create"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/delete"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/list"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/lock"
	"github.com/itera-io/taikun-cli/cmd/policyprofile/unlock"

	"github.com/spf13/cobra"
)

func NewCmdPolicyProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "policy-profile <command>",
		Short:   "Manage policy profiles",
		Aliases: []string{"pp"},
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
