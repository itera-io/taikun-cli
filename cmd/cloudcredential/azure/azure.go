package azure

import (
	"taikun-cli/cmd/cloudcredential/azure/list"

	"github.com/spf13/cobra"
)

func NewCmdAzure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure <command>",
		Short: "Manage Azure Cloud Credentials",
	}

	cmd.AddCommand(list.NewCmdList())

	return cmd
}
