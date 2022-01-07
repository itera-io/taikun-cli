package cloudcredential

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/delete"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/list"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/lock"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/unlock"

	"github.com/spf13/cobra"
)

func NewCmdCloudCredential() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cloud-credential <command>",
		Short:   "Manage Cloud Credentials",
		Aliases: []string{"cc"},
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
