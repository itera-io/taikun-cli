package unbind

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type UnbindOptions struct {
	catalogappid int32
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := cobra.Command{
		Use:   "unbind <CATALOG_APP_ID>",
		Short: "Unbind catalog app from catalog.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.catalogappid, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}

			return bindRun(cmd, &opts)
		},
	}

	return &cmd
}

func bindRun(cmd *cobra.Command, opts *UnbindOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.CatalogAppAPI.CatalogAppDelete(ctx, opts.catalogappid).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("catalog app", opts.catalogappid)

	return nil
}
