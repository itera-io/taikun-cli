package upgrade

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/spf13/cobra"
)

type UpgradeOptions struct {
	ProjectID int32
}

func NewCmdUpgrade() *cobra.Command {
	var opts UpgradeOptions

	cmd := cobra.Command{
		Use:   "upgrade <project-id>",
		Short: "Upgrade a project's version of Kubespray to the latest one",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return upgradeRun(&opts)
		},
	}

	return &cmd
}

func upgradeRun(opts *UpgradeOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := projects.NewProjectsUpgradeParams().WithV(api.Version)
	params = params.WithProjectID(opts.ProjectID)

	_, err = apiClient.Client.Projects.ProjectsUpgrade(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
