package bind

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	Flavors   []string
	ProjectID int32
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := &cobra.Command{
		Use:   "bind <project-id>",
		Short: "Bind one or multiple flavors to a project",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return bindRun(&opts)
		},
		Args: cobra.ExactArgs(1),
	}

	cmd.Flags().StringSliceVarP(&opts.Flavors, "flavors", "f", []string{}, "Flavors (required)")
	cmdutils.MarkFlagRequired(cmd, "flavors")

	return cmd
}

func bindRun(opts *BindOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.BindFlavorToProjectCommand{
		ProjectId: &opts.ProjectID,
		Flavors:   opts.Flavors,
	}
	_, response, err := myApiClient.Client.FlavorsAPI.FlavorsBindToProject(context.TODO()).BindFlavorToProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
