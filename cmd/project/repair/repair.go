package repair

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/spf13/cobra"
)

type RepairOptions struct {
	ProjectID int32
}

func NewCmdRepair() *cobra.Command {
	var opts RepairOptions

	cmd := cobra.Command{
		Use:   "repair <project-id>",
		Short: "Repair a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return repairRun(&opts)
		},
	}

	return &cmd
}

func repairRun(opts *RepairOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := projects.NewProjectsRepairParams().WithV(api.Version)
	params = params.WithProjectID(opts.ProjectID)

	_, err = apiClient.Client.Projects.ProjectsRepair(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
