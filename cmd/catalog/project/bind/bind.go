package bind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdBind() *cobra.Command {
	cmd := cobra.Command{
		Use:   "bind <CATALOG_ID> <PROJECT_ID>",
		Short: "Bind project to catalog id.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			catalogid, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}

			projectid, err := types.Atoi32(args[1])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return bindRun(catalogid, projectid)
		},
	}

	return &cmd
}

func bindRun(catalogid int32, projectid int32) (err error) {
	myApiClient := tk.NewClient()

	_, response, err := myApiClient.Client.CatalogAPI.CatalogAddProject(context.TODO(), catalogid).RequestBody([]int32{projectid}).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil

}
