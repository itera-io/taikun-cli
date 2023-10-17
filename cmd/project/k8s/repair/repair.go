package repair

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type RepairOptions struct {
	ProjectID int32
}

func NewCmdRepair() *cobra.Command {
	var opts RepairOptions

	cmd := cobra.Command{
		Use:   "repair <project-id>",
		Short: "Repair a project's Kubernetes servers",
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
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.ProjectsAPI.ProjectsRepair(context.TODO(), opts.ProjectID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := projects.NewProjectsRepairParams().WithV(taikungoclient.Version)
		params = params.WithProjectID(opts.ProjectID)

		_, err = apiClient.Client.Projects.ProjectsRepair(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
