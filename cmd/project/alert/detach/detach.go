package detach

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
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
			return detachRun(&opts)
		},
	}

	return &cmd
}

func detachRun(opts *DetachOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.AttachDetachAlertingProfileCommand{
		ProjectID: opts.ProjectID,
	}

	params := alerting_profiles.NewAlertingProfilesDetachParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.AlertingProfiles.AlertingProfilesDetach(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}
	return
}
