package list

import (
	"context"
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
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("K8S-VERSION", "kubernetesVersion"),
		field.NewVisibleWithToStringFunc("CATALOG-ID", "maintenanceModeEnabled", SetLocalCatalogId),
		field.NewVisible("CLOUD", "cloudType"),
		field.NewVisible("STATUS", "status"),
		field.NewVisible("HEALTH", "health"),
		field.NewHidden("LOCKED", "isLocked"),
		field.NewHidden("VIRTUAL", "isVirtualCluster"),
	},
)

type ListOptions struct {
	OrganizationID int32
}

// Hack, to also display the catalog id for user convenience. maintenanceModeEnabled is totally ignored
var localCatID *int32

func SetLocalCatalogId(ignore interface{}) string {
	return fmt.Sprint(*localCatID)
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <CATALOG_ID>",
		Short: "List projects bound to this catalog",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			catid, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			localCatID = &catid
			return listRun(&catid, opts.OrganizationID)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", -1, "Organization ID")
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(catid *int32, orgid int32) (err error) {
	myApiClient := tk.NewClient()

	listCommand := myApiClient.Client.CatalogAPI.CatalogList(context.TODO()).Id(*catid)
	if orgid != -1 {
		listCommand = listCommand.OrganizationId(orgid)
	}
	data, response, err := listCommand.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	if len(data.GetData()) != 1 {
		return fmt.Errorf("catalog not found")
	}
	if len(data.Data[0].BoundProjects) < 1 {
		return out.PrintResults([]interface{}{}, listFields)
	}

	return out.PrintResults(data.Data[0].BoundProjects, listFields)

}
