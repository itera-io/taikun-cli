package list

import (
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "catalogAppId"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("DESCRIPTION", "description"),
		field.NewVisibleWithToStringFunc("REPOSITORY", "repository", out.FormatRepoName),
		field.NewHidden("LOCKED", "isLocked"),
		field.NewHidden("INSTALLED-COUNT", "installedInstanceCount"),
	},
)

func NewCmdList() *cobra.Command {
	cmd := cobra.Command{
		Use:   "list <CATALOG_ID>",
		Short: "List applications bound to this catalog",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			catid, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return listRun(cmd, catid)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(cmd *cobra.Command, catid int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()

	data, response, err := myApiClient.Client.CatalogAPI.CatalogList(ctx).Id(catid).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	if len(data.GetData()) != 1 {
		return fmt.Errorf("catalog not found")
	}
	if len(data.Data[0].BoundApplications) < 1 {
		return out.PrintResults([]interface{}{}, listFields)
	}

	return out.PrintResults(data.Data[0].BoundApplications, listFields)

}
