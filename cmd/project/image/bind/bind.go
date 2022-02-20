package bind

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/images"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	ProjectID int32
	ImageIDs  []string
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := cobra.Command{
		Use:   "bind <project-id>",
		Short: "Bind one or multiple images to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return bindRun(&opts)
		},
	}

	cmd.Flags().StringSliceVarP(&opts.ImageIDs, "image-ids", "i", []string{}, "IDs of the images to bind (required)")
	cmdutils.MarkFlagRequired(&cmd, "image-ids")

	return &cmd
}

func bindRun(opts *BindOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.BindImageToProjectCommand{
		ProjectID: opts.ProjectID,
		Images:    opts.ImageIDs,
	}

	params := images.NewImagesBindImagesToProjectParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Images.ImagesBindImagesToProject(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
