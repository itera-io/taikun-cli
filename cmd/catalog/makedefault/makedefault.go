package makedefault

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

func NewCmdMakedefault() *cobra.Command {
	cmd := cobra.Command{
		Use:   "make-default <CATALOG_ID>",
		Short: "Make the catalog default for the current user.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			catalogid, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}

			return makedefaultRun(catalogid)
		},
	}

	return &cmd
}

func makedefaultRun(catalogid int32) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.CatalogMakeDefaultCommand{
		Id: &catalogid,
	}

	_, response, err := myApiClient.Client.CatalogAPI.CatalogMakeDefault(context.TODO()).CatalogMakeDefaultCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil

}
