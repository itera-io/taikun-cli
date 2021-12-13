package delete

import (
	"fmt"
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <organization-id>",
		Short: "Delete organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			organizationID, err := cmdutils.Atoi32(args[0])
			if err != nil {
				return fmt.Errorf("the given ID must be a number")
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

	params := organizations.NewOrganizationsDeleteParams().WithV(cmdutils.ApiVersion)
	params = params.WithOrganizationID(id)
	_, _, err = apiClient.Client.Organizations.OrganizationsDelete(params, apiClient)
	if err == nil {
		fmt.Println("Organization deleted")
	}

	return
}
