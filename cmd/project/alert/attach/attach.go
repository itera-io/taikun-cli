package attach

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type AttachOptions struct {
	AlertingProfileID int32
	ProjectID         int32
}

func NewCmdAttach() *cobra.Command {
	var opts AttachOptions

	cmd := cobra.Command{
		Use:   "attach <project-id>",
		Short: "Attach an alerting profile to a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ProjectID, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return attachRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.AlertingProfileID, "alerting-profile-id", "a", 0, "Alerting profile ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "alerting-profile-id")

	return &cmd
}

func attachRun(opts *AttachOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.AttachDetachAlertingProfileCommand{
		ProjectId:         &opts.ProjectID,
		AlertingProfileId: *taikuncore.NewNullableInt32(&opts.AlertingProfileID),
	}
	response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesAttach(context.TODO()).AttachDetachAlertingProfileCommand(body).Execute()
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

		body := models.AttachDetachAlertingProfileCommand{
			AlertingProfileID: opts.AlertingProfileID,
			ProjectID:         opts.ProjectID,
		}

		params := alerting_profiles.NewAlertingProfilesAttachParams().WithV(taikungoclient.Version)
		params = params.WithBody(&body)

		_, err = apiClient.Client.AlertingProfiles.AlertingProfilesAttach(params, apiClient)
		if err == nil {
			out.PrintStandardSuccess()
		}

		return
	*/
}
