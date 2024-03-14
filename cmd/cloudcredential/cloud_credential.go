package cloudcredential

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/flavors"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/images"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/list"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/lock"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/proxmox"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/remove"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/unlock"
	"github.com/spf13/cobra"
)

func NewCmdCloudCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-credential <command>",
		Short:   "Manage cloud credentials",
		Aliases: []string{"cc"},
	}

	cmd.AddCommand(aws.NewCmdAWS())
	cmd.AddCommand(azure.NewCmdAzure())
	cmd.AddCommand(flavors.NewCmdFlavors())
	//cmd.AddCommand(google.NewCmdGoogle())
	cmd.AddCommand(images.NewCmdImages())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(lock.NewCmdLock())
	cmd.AddCommand(openstack.NewCmdOpenstack())
	cmd.AddCommand(proxmox.NewCmdProxmox())
	cmd.AddCommand(remove.NewCmdDelete())
	cmd.AddCommand(unlock.NewCmdUnlock())

	return cmd
}
