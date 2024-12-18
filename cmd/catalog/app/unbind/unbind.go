package unbind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
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

			return bindRun(&opts)
		},
	}

	return &cmd
}

func bindRun(opts *UnbindOptions) (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.CatalogAppAPI.CatalogAppDelete(context.TODO(), opts.catalogappid).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("catalog app", opts.catalogappid)

	return nil
}
