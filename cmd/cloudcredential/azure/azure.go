package azure

import (
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/add"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/check"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/list"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/offers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/publishers"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/skus"
	"github.com/spf13/cobra"
)

func NewCmdAzure() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "azure <command>",
		Short: "Manage Azure cloud credentials",
	}

	cmd.AddCommand(add.NewCmdAdd())
	cmd.AddCommand(check.NewCmdCheck())
	cmd.AddCommand(list.NewCmdList())
	cmd.AddCommand(offers.NewCmdOffers())
	cmd.AddCommand(publishers.NewCmdPublishers())
	cmd.AddCommand(skus.NewCmdSKUs())

	return cmd
}
