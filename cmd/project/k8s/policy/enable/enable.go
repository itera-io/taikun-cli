package enable

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
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
	body := taikuncore.EnableGatekeeperCommand{
		ProjectId:    &opts.ProjectID,
		OpaProfileId: &opts.PolicyProfileID,
	}
	response, err := myApiClient.Client.OpaProfilesAPI.OpaprofilesEnableGatekeeper(context.TODO()).EnableGatekeeperCommand(body).Execute()
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

		body := models.EnableGatekeeperCommand{
			OpaProfileID: opts.PolicyProfileID,
			ProjectID:    opts.ProjectID,
		}

		params := opa_profiles.NewOpaProfilesEnableGatekeeperParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.OpaProfiles.OpaProfilesEnableGatekeeper(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
