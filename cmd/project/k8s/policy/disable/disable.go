package disable

import (
	"context"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DisableOptions struct {
	ProjectID int32
}

func NewCmdDisable() *cobra.Command {
	var opts DisableOptions

	cmd := cobra.Command{
		Use:   "disable <project-id>",
		Short: "Disable policy for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return disableRun(&opts)
		},
	}

	return &cmd
}

func disableRun(opts *DisableOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.DeploymentDisableOpaCommand{
		ProjectId: &opts.ProjectID,
	}
	_, response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentDisableOpa(context.TODO()).DeploymentDisableOpaCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return

}
