package delete

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {

	cmd := cobra.Command{
		Use:   "delete <CATALOG_ID>",
		Short: "Delete a catalog",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			catalogid, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return deletecatalogRun(catalogid)
		},
	}

	return &cmd
}

func deletecatalogRun(catalogid int32) (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.CatalogAPI.CatalogDelete(context.TODO(), catalogid).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("catalog", catalogid)

	return nil
}
