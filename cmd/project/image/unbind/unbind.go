package unbind

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
			return unbindRun(cmd, &opts)
		},
	}

	return &cmd
}

func unbindRun(cmd *cobra.Command, opts *UnbindOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepere the arguments for the query
	body := taikuncore.DeleteImageFromProjectCommand{
		Ids: opts.ImageBindingIDs,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.ImagesAPI.ImagesUnbindImagesFromProject(ctx).DeleteImageFromProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
