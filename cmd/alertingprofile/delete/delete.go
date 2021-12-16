package delete

import (
	"taikun-cli/api"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <alerting-profile-id>",
		Short: "Delete an alerting profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := utils.Atoi32(args[0])
			if err != nil {
				return utils.WrongIDArgumentFormatError
			}
			return deleteRun(id)
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

	params := alerting_profiles.NewAlertingProfilesDeleteParams().WithV(utils.ApiVersion).WithBody(&body)
	_, _, err = apiClient.Client.AlertingProfiles.AlertingProfilesDelete(params, apiClient)
	if err == nil {
		utils.PrintDeleteSuccess("Alerting profile", id)
	}

	return
}
