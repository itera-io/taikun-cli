package bind

import (
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

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
		Use:   "bind <flavor-name>...",
		Short: "Bind one or multiple flavors to a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Flavors = args
			return bindRun(&opts)
		},
		Args: cobra.MinimumNArgs(1),
	}

	cmd.Flags().Int32VarP(&opts.ProjectID, "project-id", "p", 0, "Project ID (required)")
	cmdutils.MarkFlagRequired(cmd, "project-id")

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
	params := flavors.NewFlavorsBindToProjectParams().WithV(cmdutils.ApiVersion).WithBody(&body)
	response, err := apiClient.Client.Flavors.FlavorsBindToProject(params, apiClient)
	if err == nil {
		cmdutils.PrettyPrint(response.Payload)
	}

	return
}
