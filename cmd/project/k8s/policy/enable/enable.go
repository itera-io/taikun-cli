package enable

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type EnableOptions struct {
	ProjectID       int32
	PolicyProfileID int32
}

func NewCmdEnable() *cobra.Command {
	var opts EnableOptions

	cmd := cobra.Command{
		Use:   "enable <project-id>",
		Short: "Enable policy for a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return enableRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.PolicyProfileID, "policy-profile-id", "p", 0, "Policy profile ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "policy-profile-id")

	return &cmd
}

func enableRun(opts *EnableOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.DeploymentOpaEnableCommand{
		ProjectId:       &opts.ProjectID,
		OpaCredentialId: &opts.PolicyProfileID,
	}
	response, err := myApiClient.Client.ProjectDeploymentAPI.ProjectDeploymentEnableOpa(context.TODO()).DeploymentOpaEnableCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
