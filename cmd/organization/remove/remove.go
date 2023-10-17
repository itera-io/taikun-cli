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
		Use:   "delete <organization-id>...",
		Short: "Delete one or more organizations",
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

func deleteRun(orgID int32) (err error) {
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.OrganizationsAPI.OrganizationsDelete(context.TODO(), orgID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("Organization", orgID)
	return

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := organizations.NewOrganizationsDeleteParams().WithV(taikungoclient.Version)
		params = params.WithOrganizationID(orgID)

		_, _, err = apiClient.Client.Organizations.OrganizationsDelete(params, apiClient)
		if err == nil {
			out.PrintDeleteSuccess("Organization", orgID)
		}

		return
	*/
}
