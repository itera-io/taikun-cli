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
	ProjectID int32
}

func NewCmdDelete() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "delete <project-id>",
		Short: "Delete one or more virtual projects",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			optsList := make([]*DeleteOptions, len(args))
			for i, arg := range args {
				projectID, err := types.Atoi32(arg)
				if err != nil {
					return cmderr.ErrIDArgumentNotANumber
				}
				optsList[i] = &DeleteOptions{
					ProjectID: projectID,
				}
			}
			return deleteMultiple(optsList)
		},
		Aliases: cmdutils.DeleteAliases,
	}

	return cmd
}

func deleteMultiple(optsList []*DeleteOptions) error {
	errorOccured := false

	for _, opts := range optsList {
		if err := deleteRun(opts); err != nil {
			errorOccured = true

			_, _ = fmt.Fprintln(os.Stderr, err)
		}
	}

	if errorOccured {
		_, _ = fmt.Fprintln(os.Stderr)
		return errors.New("Failed to delete one or more virtual projects")
	}

	return nil
}

func deleteRun(opts *DeleteOptions) (err error) {
	myApiClient := tk.NewClient()
	body := taikuncore.DeleteVirtualClusterCommand{
		ProjectId: &opts.ProjectID,
	}

	request, err := myApiClient.Client.VirtualClusterAPI.VirtualClusterDelete(context.TODO()).DeleteVirtualClusterCommand(body).Execute()
	if err != nil {
		return tk.CreateError(request, err)
	}
	out.PrintDeleteSuccess("Virtual Project", opts.ProjectID)

	return
}
