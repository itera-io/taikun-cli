package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByAndReverseFlags(cmd, models.UserForListDto{})
	cmdutils.AddLimitFlag(cmd)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	users, err := ListUsers(opts)
	if err != nil {
		return err
	}

	out.PrintResults(users,
		"id",
		"username",
		"role",
		"organizationName",
		"email",
	)
	return
}

func ListUsers(opts *ListOptions) (userList []*models.UserForListDto, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	userList = make([]*models.UserForListDto, 0)
	for {
		response, err := apiClient.Client.Users.UsersList(params, apiClient)
		if err != nil {
			return nil, err
		}
		userList = append(userList, response.Payload.Data...)
		usersCount := int32(len(userList))
		if config.Limit != 0 && usersCount >= config.Limit {
			break
		}
		if usersCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&usersCount)
	}

	if config.Limit != 0 && int32(len(userList)) > config.Limit {
		userList = userList[:config.Limit]
	}

	return
}
