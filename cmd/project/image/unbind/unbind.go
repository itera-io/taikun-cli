package unbind

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/client/images"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type UnbindOptions struct {
	ImageBindingIDs []int32
}

func NewCmdUnbind() *cobra.Command {
	var opts UnbindOptions

	cmd := cobra.Command{
		Use:   "unbind <image-binding-id>...",
		Short: "Unbind one or more images from a project",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ImageBindingIDs, err = cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return unbindRun(&opts)
		},
	}

	return &cmd
}

func unbindRun(opts *UnbindOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteImageFromProjectCommand{
		Ids: opts.ImageBindingIDs,
	}

	params := images.NewImagesUnbindImagesFromProjectParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Images.ImagesUnbindImagesFromProject(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
