package unbind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdUnbind() *cobra.Command {
	cmd := cobra.Command{
		Use:   "unbind <CATALOG_ID> <PROJECT_ID>",
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
			return unbindRun(catalogid, projectid)
		},
	}

	return &cmd
}

func unbindRun(catalogid int32, projectid int32) (err error) {
	myApiClient := tk.NewClient()

	response, err := myApiClient.Client.CatalogAPI.CatalogDeleteProject(context.TODO(), catalogid).RequestBody([]int32{projectid}).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil

}
