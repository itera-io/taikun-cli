package bind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepere the arguments for the query
	body := taikuncore.BindImageToProjectCommand{
		ProjectId: &opts.ProjectID,
		Images:    opts.ImageIDs,
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.Client.ImagesAPI.ImagesBindImagesToProject(context.TODO()).BindImageToProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
