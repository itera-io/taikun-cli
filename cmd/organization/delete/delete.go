package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/utils/format"
	"taikun-cli/utils/types"

	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <organization-id>",
		Short: "Delete organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			organizationID, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.WrongIDArgumentFormatError
			}
			return deleteRun(organizationID)
		},
	}

	return cmd
}

func deleteRun(id int32) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := organizations.NewOrganizationsDeleteParams().WithV(apiconfig.Version)
	params = params.WithOrganizationID(id)
	_, _, err = apiClient.Client.Organizations.OrganizationsDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Organization", id)
	}

	return
}
