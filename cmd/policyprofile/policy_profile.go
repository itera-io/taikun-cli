package policyprofile

import (
	"taikun-cli/cmd/policyprofile/create"
	"taikun-cli/cmd/policyprofile/delete"
	"taikun-cli/cmd/policyprofile/list"

	"github.com/spf13/cobra"
)

func NewCmdPolicyProfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "policy-profile <command>",
		Short: "Manage policy profiles",
	}

	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(delete.NewCmdDelete())

	return cmd
}
