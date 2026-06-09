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
		Use:   "delete <access-profile-id>...",
		Short: "Delete one or more access profiles",
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

func deleteRun(ctx context.Context, accessProfileID int32) (err error) {
	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.AccessProfilesAPI.AccessprofilesDelete(ctx, accessProfileID).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}
	out.PrintDeleteSuccess("Access profile", accessProfileID)
	return
}
