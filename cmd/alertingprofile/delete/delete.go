package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <alerting-profile-id>...",
		Short: "Delete one or more alerting profiles",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.DeleteAlertingProfilesCommand{ID: id}

	params := alerting_profiles.NewAlertingProfilesDeleteParams().WithV(apiconfig.Version).WithBody(&body)
	_, _, err = apiClient.Client.AlertingProfiles.AlertingProfilesDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Alerting profile", id)
	}

	return
}
