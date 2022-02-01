package bind

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/flavors"
	"github.com/itera-io/taikungoclient/models"
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
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.BindFlavorToProjectCommand{
		ProjectID: opts.ProjectID,
		Flavors:   opts.Flavors,
	}
	params := flavors.NewFlavorsBindToProjectParams().WithV(api.Version).WithBody(&body)
	_, err = apiClient.Client.Flavors.FlavorsBindToProject(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
