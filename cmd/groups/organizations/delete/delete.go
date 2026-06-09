package delete

import (
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
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
			return deleteOrganizationsFromGroup(cmd, groupID, &opts)
		},
	}

	cmd.Flags().Int32SliceVarP(&opts.OrganizationIDs, "organization-id", "O", nil, "Organization IDs")
	return &cmd
}

func deleteOrganizationsFromGroup(cmd *cobra.Command, groupID int32, opts *DeleteOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// input parameters sanity check
	if len(opts.OrganizationIDs) == 0 {
		return fmt.Errorf("no organization IDs are specified")
	}
	myApiClient := tk.NewClient()

	body := *taikuncore.NewDeleteOrganizationFromGroupCommand(groupID, opts.OrganizationIDs)
	response, err := myApiClient.Client.GroupsAPI.GroupsDeleteOrganizations(ctx).DeleteOrganizationFromGroupCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
