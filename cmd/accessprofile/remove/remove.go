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
		Use:   "delete <access-profile-id>...",
		Short: "Delete one or more access profiles",
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

func deleteRun(accessProfileID int32) (err error) {
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.AccessProfilesAPI.AccessprofilesDelete(context.TODO(), accessProfileID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("Access profile", accessProfileID)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := access_profiles.NewAccessProfilesDeleteParams().WithV(taikungoclient.Version).WithID(accessProfileID)

		_, _, err = apiClient.Client.AccessProfiles.AccessProfilesDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Access profile", accessProfileID)
		}

		return
	*/
}
