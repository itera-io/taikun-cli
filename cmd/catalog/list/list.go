package list

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization/list"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisibleWithToStringFunc("ORG", "organizationId", formatOrganizationName),
		field.NewVisible("DEFAULT", "isDefault"),
		field.NewVisible("DESCRIPTION", "description"),
		field.NewVisible("LOCKED", "isLocked"),
	},
)

type ListOptions struct {
	Organization int32
	Limit        int32
}

// We want do display organization name, but want to download the organizations only once.
var organizationCache map[int32]string

func initOrganizationCache() {
	if organizationCache != nil {
		return
	}

	organizationCache = make(map[int32]string)
	var opts list.ListOptions
	organizations, err := list.ListOrganizations(&opts)
	if err != nil {
		fmt.Println("Error fetching organizations:", err)
		return
	}

	for _, org := range organizations {
		organizationCache[*org.Id] = *org.Name.Get()
	}
}

func formatOrganizationName(orgidinput interface{}) string {
	initOrganizationCache()

	orgid := int32(orgidinput.(float64))
	if name, found := organizationCache[orgid]; found {
		return name
	}

	return ""
}

// New command
func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List available catalogs",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.Organization, "organization", "o", 0, "Id of the organization to use for the list-private (partner)")
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "catalogs", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	var catalogs = make([]taikuncore.CatalogListDto, 0)

	myRequest := myApiClient.Client.CatalogAPI.CatalogList(context.TODO())
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	if opts.Organization != 0 {
		myRequest = myRequest.OrganizationId(opts.Organization)
	}

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		catalogs = append(catalogs, data.GetData()...)
		applicationsCount := int32(len(catalogs))

		if opts.Limit != 0 && applicationsCount >= opts.Limit {
			break
		}

		if applicationsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(applicationsCount)
	}
	if opts.Limit != 0 && int32(len(catalogs)) > opts.Limit {
		catalogs = catalogs[:opts.Limit]
	}

	return out.PrintResults(catalogs, listFields)

}
