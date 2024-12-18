package bind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("REPO-NAME", "appRepoName"),
		field.NewVisible("REPO-ORG", "appRepoOrganizationName"),
		field.NewVisible("CATALOG-NAME", "catalogName"),
		field.NewVisible("CATALOG-ID", "catalogId"),
		field.NewVisible("DESCRIPTION", "description"),
		field.NewVisible("APP-VERSION", "appVersion"),
	},
)

type BindOptions struct {
	catalogid int32
	repoName  string
	appName   string
	version   string
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := cobra.Command{
		Use:   "bind <CATALOG_ID> <APP_NAME> <REPOSITORY_NAME>",
		Short: "Bind project to catalog id.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.catalogid, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}

			opts.appName = args[1]
			opts.repoName = args[2]

			return bindRun(&opts)
		},
	}

	return &cmd
}

func bindRun(opts *BindOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.CreateCatalogAppCommand{
		RepoName:    *taikuncore.NewNullableString(&opts.repoName),
		PackageName: *taikuncore.NewNullableString(&opts.appName),
		CatalogId:   &opts.catalogid,
		Parameters:  make([]taikuncore.CatalogAppParamsDto, 0),
	}
	if opts.version != "" {
		body.SetVersion(opts.version)
	}

	data, response, err := myApiClient.Client.CatalogAppAPI.CatalogAppCreate(context.TODO()).CreateCatalogAppCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResult(data, addFields)
}
