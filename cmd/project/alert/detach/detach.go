package detach

import (
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

type DetachOptions struct {
	ProjectID int32
}

func NewCmdDetach() *cobra.Command {
	var opts DetachOptions

	cmd := cobra.Command{
		Use:   "detach <project-id>",
		Short: "Detach a project's alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return detachRun(cmd, &opts)
		},
	}

	return &cmd
}

func detachRun(cmd *cobra.Command, opts *DetachOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	myApiClient := tk.NewClient()
	body := taikuncore.AttachDetachAlertingProfileCommand{
		ProjectId: &opts.ProjectID,
	}
	response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesDetach(ctx).AttachDetachAlertingProfileCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintStandardSuccess()
	return

}
