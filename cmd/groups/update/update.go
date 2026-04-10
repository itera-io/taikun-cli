package update

import (
	"context"
	"fmt"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UpdateOptions struct {
	GroupName  string
	ClaimValue string
}

func NewCmdUpdateGroup() *cobra.Command {
	var opts UpdateOptions

	cmd := cobra.Command{
		Use:   "update <GROUP_ID>",
		Short: "Update existing group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return updateGroup(groupID, &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.GroupName, "name", "n", "", "Group name")
	cmd.Flags().StringVarP(&opts.ClaimValue, "claim-value", "v", "false", "Claim value")

	return &cmd
}

func updateGroup(groupID int32, opts *UpdateOptions) (err error) {
	// input parameters sanity check
	if opts.GroupName == "" && opts.ClaimValue == "" {
		return fmt.Errorf("no parameters to update are passed")
	}
	myApiClient := tk.NewClient()

	body := taikuncore.UpdateGroupDto{
		Name:       *taikuncore.NewNullableString(&opts.GroupName),
		ClaimValue: *taikuncore.NewNullableString(&opts.ClaimValue),
	}

	response, err := myApiClient.Client.GroupsAPI.GroupsUpdate(context.TODO(), groupID).UpdateGroupDto(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
