package cloudcredential

import (
	"taikun-cli/cmd/cloudcredential/aws"
	"taikun-cli/cmd/cloudcredential/azure"
	"taikun-cli/cmd/cloudcredential/delete"
	"taikun-cli/cmd/cloudcredential/list"
	"taikun-cli/cmd/cloudcredential/lock"
	"taikun-cli/cmd/cloudcredential/openstack"
	"taikun-cli/cmd/cloudcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdBillingCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-credential <command>",
		Short:   "Manage Cloud Credentials",
		Aliases: []string{"cloud", "cc"},
	}

	cmd.AddCommand(aws.NewCmdAWS())
	cmd.AddCommand(azure.NewCmdAzure())
	cmd.AddCommand(delete.NewCmdDelete())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(openstack.NewCmdOpenstack())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
