package lock

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
)

func NewCmdLock() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lock <catalog-id>",
		Short: "Lock a catalog",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return lockRun(cmd, id)
		},
	}

	return cmd
}

func lockRun(cmd *cobra.Command, catalogID int32) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CatalogLockManagementCommand{
		Id:   &catalogID,
		Mode: *taikuncore.NewNullableString(&types.LockedMode),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.CatalogAPI.CatalogLock(ctx).CatalogLockManagementCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return

}
