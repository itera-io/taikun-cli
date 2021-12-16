package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <access-profile-id>",
		Short: "Delete an access profile",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
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

	params := access_profiles.NewAccessProfilesDeleteParams().WithV(apiconfig.Version).WithID(id)
	_, _, err = apiClient.Client.AccessProfiles.AccessProfilesDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Access profile", id)
	}

	return
}
