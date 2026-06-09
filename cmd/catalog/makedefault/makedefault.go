package makedefault

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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

			return makedefaultRun(cmd, catalogid)
		},
	}

	return &cmd
}

func makedefaultRun(cmd *cobra.Command, catalogid int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	body := taikuncore.CatalogMakeDefaultCommand{
		Id: &catalogid,
	}

	response, err := myApiClient.Client.CatalogAPI.CatalogMakeDefault(ctx).CatalogMakeDefaultCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil

}
