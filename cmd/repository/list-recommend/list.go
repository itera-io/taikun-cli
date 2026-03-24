package list_recommend

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewHidden("ID", "repositoryId"),
		field.NewVisible("REPOSITORY-NAME", "name"),
		field.NewVisible("REPOSITORY-ORG", "organizationName"),
		field.NewVisible("URL", "url"),
		field.NewVisibleWithToStringFunc("BOUND", "isBound", out.FormatDisabled),
		field.NewVisible("HAS-APP", "hasCatalogApp"),
	},
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list-recommend",
		Short: "List available managed repositories",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	//cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}

	myApiClient := tk.NewClient()
	var repositories = make([]taikuncore.ArtifactRepositoryDto, 0)

	// Recommended
	myquery := myApiClient.Client.AppRepositoriesAPI.RepositoryRecommendedList(context.TODO())
	if orgID != 0 {
		myquery = myquery.OrganizationId(orgID)
	}
	data, response, err := myquery.Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	repositories = append(repositories, data...)

	return out.PrintResults(repositories, listFields)

}
