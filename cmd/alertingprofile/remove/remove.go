package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
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
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultiple(ids, deleteRun)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(alertingProfileID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesDelete(context.TODO(), alertingProfileID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintDeleteSuccess("Alerting profile", alertingProfileID)
	return

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.DeleteAlertingProfilesCommand{ID: alertingProfileID}

		params := alerting_profiles.NewAlertingProfilesDeleteParams().WithV(taikungoclient.Version).WithBody(&body)

		_, _, err = apiClient.Client.AlertingProfiles.AlertingProfilesDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Alerting profile", alertingProfileID)
		}

		return
	*/
}
