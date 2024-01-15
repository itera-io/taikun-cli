package list

import (
	"context"
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
		field.NewVisible(
			"BINDING-ID", "id",
		),
		field.NewVisible(
			"IMAGE", "name",
		),
		field.NewVisible(
			"IMAGE-ID", "imageId",
		),
		field.NewVisible(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
		field.NewVisible(
			"SIZE", "size",
		),
		field.NewHidden(
			"CLOUD-ID", "cloudId",
		),
		field.NewHidden(
			"WINDOWS", "isWindows",
		),
		field.NewHidden(
			"AWS", "isAws",
		),
		field.NewHidden(
			"AZURE", "isAzure",
		),
		field.NewHidden(
			"OPENSTACK", "isOpenstack",
		),
	},
)

type ListOptions struct {
	ProjectID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <project-id>",
		Short: "List a project's bound images",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return listRun(&opts)
		},
	}

	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.ImagesAPI.ImagesSelectedImagesForProject(context.TODO()).ProjectId(opts.ProjectID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResults(data.GetData(), listFields)

}
