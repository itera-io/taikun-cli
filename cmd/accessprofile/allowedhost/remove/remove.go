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
		Use:   "delete <allowed-host-id>...",
		Short: "Delete an allowed host",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := cmdutils.APIContext(cmd)
			defer cancel()
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return cmdutils.DeleteMultiple(ids, func(id int32) error {
				return deleteRun(ctx, id)
			})
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteRun(ctx context.Context, allowedHostID int32) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.AllowedHostAPI.AllowedhostDelete(ctx, allowedHostID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("Allowed host", allowedHostID)
	return

}
