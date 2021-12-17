package delete

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/organizations"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <organization-id>...",
		Short: "Delete one or more organizations",
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

	params := organizations.NewOrganizationsDeleteParams().WithV(apiconfig.Version)
	params = params.WithOrganizationID(id)
	_, _, err = apiClient.Client.Organizations.OrganizationsDelete(params, apiClient)
	if err == nil {
		format.PrintDeleteSuccess("Organization", id)
	}

	return
}
