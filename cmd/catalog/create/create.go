package create

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Description    string
	OrganizationID int32
}

func NewCmdCreatecatalog() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <CATALOG_NAME>",
		Short: "Create a new catalog",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return createcatalogRun(args[0], opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmd.Flags().StringVarP(&opts.Description, "description", "d", "", "Description (min 3 characters)")
	_ = cmd.MarkFlagRequired("description")

	return &cmd
}

func createcatalogRun(catalogname string, opts CreateOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.CreateCatalogCommand{
		Name:        *taikuncore.NewNullableString(&catalogname),
		Description: *taikuncore.NewNullableString(&opts.Description),
	}

	if opts.OrganizationID != 0 {
		body.SetOrganizationId(opts.OrganizationID)
	}

	_, response, err := myApiClient.Client.CatalogAPI.CatalogCreate(context.TODO()).CreateCatalogCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()

	return nil
}
