package remove

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	"github.com/spf13/cobra"
)

func NewCmdDelete() *cobra.Command {
	var id int32
	cmd := &cobra.Command{
		Use:   "delete <allowed-host-id>...",
		Short: "Delete an allowed host",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err = types.Atoi32(args[0])
			if err != nil {
				return
			}
			return deleteRun(id)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(allowedHostID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.AllowedHostAPI.AllowedhostDelete(context.TODO(), allowedHostID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("Allowed host", allowedHostID)
	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}
		params := allowed_host.NewAllowedHostDeleteParams().WithV(taikungoclient.Version).WithID(allowedHostID)
		_, _, err = apiClient.Client.AllowedHost.AllowedHostDelete(params, apiClient)

		if err == nil {
			out.PrintDeleteSuccess("Allowed host", allowedHostID)
		}

		return
	*/
}
