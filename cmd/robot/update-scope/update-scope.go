package updatescope

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UpdateScopeOptions struct {
	Scopes []string
}

func NewCmdUpdateScope() *cobra.Command {
	var opts UpdateScopeOptions

	cmd := cobra.Command{
		Use:   "update-scope <ROBOT_ID>",
		Short: "Update robot user scope",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return updateRobotScope(args[0], &opts)
		},
	}

	cmd.Flags().StringSliceVarP(&opts.Scopes, "scope", "s", nil, "Scope of the robot user")
	_ = cmd.MarkFlagRequired("scope")

	return &cmd
}

func updateRobotScope(robotID string, opts *UpdateScopeOptions) (err error) {
	myApiClient := tk.NewClient()

	body := taikuncore.UpdateRobotScopeCommand{
		Id:     *taikuncore.NewNullableString(&robotID),
		Scopes: opts.Scopes,
	}

	response, err := myApiClient.Client.RobotAPI.RobotUpdateScope(context.TODO()).UpdateRobotScopeCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
