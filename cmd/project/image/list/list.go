package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/images"
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := images.NewImagesGetSelectedImagesForProjectParams().WithV(api.Version)
	params = params.WithProjectID(&opts.ProjectID)

	response, err := apiClient.Client.Images.ImagesGetSelectedImagesForProject(params, apiClient)
	if err == nil {
		return out.PrintResults(response.Payload.Data, listFields)
	}

	return
}
