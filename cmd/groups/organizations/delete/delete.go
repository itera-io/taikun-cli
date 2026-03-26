package delete

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type DeleteOptions struct {
	OrganizationIDs []int32
}

func NewCmdDeleteOrganizations() *cobra.Command {
	opts := DeleteOptions{
		OrganizationIDs: make([]int32, 0),
	}

	cmd := cobra.Command{
		Use:   "delete <GROUP_ID>",
		Short: "Remove organizations from the group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return deleteOrganizationsFromGroup(groupID, &opts)
		},
	}

	cmd.Flags().Int32SliceVarP(&opts.OrganizationIDs, "organization-id", "O", nil, "Organization IDs")
	return &cmd
}

func deleteOrganizationsFromGroup(groupID int32, opts *DeleteOptions) (err error) {
	// input parameters sanity check
	if len(opts.OrganizationIDs) == 0 {
		return fmt.Errorf("no organization IDs are specified")
	}
	myApiClient := tk.NewClient()

	body := *taikuncore.NewDeleteOrganizationFromGroupCommand(groupID, opts.OrganizationIDs)
	response, err := myApiClient.Client.GroupsAPI.GroupsDeleteOrganizations(context.TODO()).DeleteOrganizationFromGroupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
