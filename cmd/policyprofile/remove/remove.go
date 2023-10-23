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
		Use:   "delete <policy-profile-id>...",
		Short: "Delete one or more policy profiles",
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

func deleteRun(policyProfileID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.OpaProfilesAPI.OpaprofilesDelete(context.TODO(), policyProfileID).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}

	out.PrintDeleteSuccess("Policy profile", policyProfileID)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := &models.DeleteOpaProfileCommand{ID: policyProfileID}
		params := opa_profiles.NewOpaProfilesDeleteParams().WithV(taikungoclient.Version).WithBody(body)

		_, err = apiClient.Client.OpaProfiles.OpaProfilesDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Policy profile", policyProfileID)
		}

		return
	*/
}
