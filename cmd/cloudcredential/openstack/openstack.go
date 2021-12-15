package openstack

import (
	"taikun-cli/cmd/cloudcredential/openstack/list"

	"github.com/spf13/cobra"
)

func NewCmdOpenstack() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "openstack <command>",
		Short: "Manage OpenStack Cloud Credentials",
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
