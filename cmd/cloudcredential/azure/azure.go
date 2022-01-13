package azure

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/create"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/list"

	"github.com/spf13/cobra"
)

func NewCmdAzure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure <command>",
		Short: "Manage Azure cloud credentials",
	}

	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(create.NewCmdCreate())
	cmd.AddCommand(list.NewCmdList())

	return cmd
}
