package add

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	OrganizationID int32
	Role           string
	Projects       []int32
}

func NewCmdAddOrganizations() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <GROUP_ID>",
		Short: "Add organization to the group",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			groupID, err := types.Atoi32(args[0])
			if err != nil {
				return err
			}
			return addOrganizationToGroup(groupID, &opts)
		},
	}

	cmdutils.AddOrgIDFlag(&cmd, &opts.OrganizationID)
	cmd.Flags().StringVarP(&opts.Role, "role", "r", "", "Role to add to the group")
	opts.Projects = *cmd.Flags().Int32SliceP("project-id", "p", nil, "Project ID")

	_ = cmd.MarkFlagRequired("organization-id")
	_ = cmd.MarkFlagRequired("role")
	return &cmd
}

func addOrganizationToGroup(groupID int32, opts *AddOptions) (err error) {
	role, err := taikuncore.NewAccessLevelRolesFromValue(opts.Role)
	if err != nil {
		return err
	}
	orgDto := *taikuncore.NewCreateGroupOrganizationDto(opts.OrganizationID, *role)
	orgDto.SetProjects(opts.Projects)
	orgDtos := []taikuncore.CreateGroupOrganizationDto{orgDto}

	myApiClient := tk.NewClient()
	response, err := myApiClient.Client.GroupsAPI.GroupsAddOrganizations(context.TODO(), groupID).CreateGroupOrganizationDto(orgDtos).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil
}
