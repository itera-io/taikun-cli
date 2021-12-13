package delete

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <access-profile-id>",
		Short: "Delete an access profile",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("accepts 1 arg, received %d", len(args))
			}
			_, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given id must be a number")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			id, _ := cmdutils.Atoi32(args[0])
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

	params := access_profiles.NewAccessProfilesDeleteParams().WithV(cmdutils.ApiVersion).WithID(id)
	_, _, err = apiClient.Client.AccessProfiles.AccessProfilesDelete(params, apiClient)
	if err == nil {
		fmt.Println("Access Profile deleted")
	}

	return
}
