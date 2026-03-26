package update

import (
	"context"

	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type UpdateOptions struct {
	OrganizationID int32
	Role           string
	Projects       []int32
}

func NewCmdUpdateOrganization() *cobra.Command {
	var opts UpdateOptions

	cmd := cobra.Command{
		Use:   "update <GROUP_ID>",
		Short: "Update organization within group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return updateOrganizationToGroup(groupID, &opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "O", 0, "ID of organization to add to the group")
	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role to add to the group")
	opts.Projects = *cmd.Flags().Int32SliceP("project-id", "p", nil, "Project ID")

	_ = cmd.MarkFlagRequired("organization-id")
	_ = cmd.MarkFlagRequired("role")
	return &cmd
}

func updateOrganizationToGroup(groupID int32, opts *UpdateOptions) (err error) {
	orgDto := *taikuncore.NewUpdateGroupOrganizationDto(opts.Role)
	orgDto.SetProjects(opts.Projects)

	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.GroupsAPI.GroupsUpdateGroupOrganization(context.TODO(), groupID, opts.OrganizationID).UpdateGroupOrganizationDto(orgDto).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
