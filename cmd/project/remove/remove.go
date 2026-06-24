package remove

import (
	"context"
	"errors"
	"fmt"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"os"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	Force     bool
	ProjectID int32
}

func NewCmdDelete() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "delete <project-id>...",
		Short: "Delete one or more projects",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, cancel := cmdutils.APIContext(cmd)
			defer cancel()
			optsList := make([]*DeleteOptions, len(args))
			for i, arg := range args {
				projectID, err := types.Atoi32(arg)
				if err != nil {
					return cmderr.ErrIDArgumentNotANumber
				}
				optsList[i] = &DeleteOptions{
					Force:     force,
					ProjectID: projectID,
				}
			}
			return deleteMultiple(ctx, optsList)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force delete")

	return cmd
}

func deleteMultiple(ctx context.Context, optsList []*DeleteOptions) error {
	errorOccured := false

	for _, opts := range optsList {
		if err := deleteRun(ctx, opts); err != nil {
			errorOccured = true

			fmt.Fprintln(os.Stderr, err)
		}
	}

	if errorOccured {
		fmt.Fprintln(os.Stderr)
		return errors.New("failed to delete one or more projects")
	}

	return nil
}

func deleteRun(ctx context.Context, opts *DeleteOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.DeleteProjectCommand{
		ProjectId:     &opts.ProjectID,
		IsForceDelete: &opts.Force,
	}
	request, err := myApiClient.Client.ProjectsAPI.ProjectsDelete(ctx).DeleteProjectCommand(body).Execute()
	if err != nil {
		return tk.CreateError(request, err)
	}
	out.PrintDeleteSuccess("Project", opts.ProjectID)

	return
}
