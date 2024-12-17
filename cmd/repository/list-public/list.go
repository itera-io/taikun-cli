package list_public

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
		field.NewHidden("ID", "repositoryId"),
		field.NewVisible("REPOSITORY-NAME", "name"),
		field.NewVisible("REPOSITORY-ORG", "organizationName"),
		field.NewVisible("URL", "url"),
		field.NewVisibleWithToStringFunc("BOUND", "isBound", out.FormatDisabled),
		field.NewVisible("HAS-APP", "hasCatalogApp"),
	},
)

type ListOptions struct {
	Organization int32
	Limit        int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list-public",
		Short: "List available public repositories",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.Organization, "organization", "o", 0, "Id of the organization to use for the list-public (partner) - default is your home organization.")
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "repositorypublic", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	var repositories = make([]taikuncore.ArtifactRepositoryDto, 0)

	// Public
	myRequest := myApiClient.Client.AppRepositoriesAPI.RepositoryAvailableList(context.TODO()).IsPrivate(false)
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

		repositories = append(repositories, data.GetData()...)
		applicationsCount := int32(len(repositories))

		if opts.Limit != 0 && applicationsCount >= opts.Limit {
			break
		}

		if applicationsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(applicationsCount)
	}
	if opts.Limit != 0 && int32(len(repositories)) > opts.Limit {
		repositories = repositories[:opts.Limit]
	}

	return out.PrintResults(repositories, listFields)

}
