package project

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	Force bool
}

func NewCmdVms() *cobra.Command {
	var opts AddOptions
	cmd := cobra.Command{
		Use:   "project [project-id]",
		Short: "Prepare project for migrate",
		Long:  "Prepare project for migrate",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return myMigrateRun(id, &opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.Force, "force", "f", false, "Do the migration even though it will create resources in the project (repair)")

	return &cmd
}

func myMigrateRun(projectId int32, opts *AddOptions) (err error) {
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.TofuMigrateCommand{
		ProjectId: &projectId,
		Force:     &opts.Force,
	}

	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentTofuMigrate(context.TODO()).TofuMigrateCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return
}
