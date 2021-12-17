package policyprofile

import (
	"taikun-cli/cmd/policyprofile/create"
	"taikun-cli/cmd/policyprofile/delete"
	"taikun-cli/cmd/policyprofile/list"
	"taikun-cli/cmd/policyprofile/lock"
	"taikun-cli/cmd/policyprofile/unlock"

	"github.com/spf13/cobra"
)

func NewCmdPolicyProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy-profile <command>",
		Short: "Manage policy profiles",
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
