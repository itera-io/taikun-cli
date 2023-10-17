package disable

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/utils/types"
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
	body := taikuncore.DisableGatekeeperCommand{
		ProjectId: &opts.ProjectID,
	}
	response, err := myApiClient.Client.OpaProfilesAPI.OpaprofilesDisableGatekeeper(context.TODO()).DisableGatekeeperCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.DisableGatekeeperCommand{
			ProjectID: opts.ProjectID,
		}

		params := opa_profiles.NewOpaProfilesDisableGatekeeperParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.OpaProfiles.OpaProfilesDisableGatekeeper(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
