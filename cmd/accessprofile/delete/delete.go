package delete

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	ID int32
}

func NewCmdDelete() *cobra.Command {
	var opts DeleteOptions

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an access profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			return deleteRun(&opts)
		},
	}

	cmd.Flags().Int32Var(&opts.ID, "id", 0, "ID (required)")
	cmdutils.MarkFlagRequired(cmd, "id")

	return cmd
}

func deleteRun(opts *DeleteOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := access_profiles.NewAccessProfilesDeleteParams().WithV(cmdutils.ApiVersion).WithID(opts.ID)
	_, _, err = apiClient.Client.AccessProfiles.AccessProfilesDelete(params, apiClient)
	if err == nil {
		fmt.Println("Access Profile deleted")
	}

	return
}
