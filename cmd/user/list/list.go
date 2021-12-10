package list

import (
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
	// TODO add other flags
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	// TODO add other flags

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersListParams().WithV(cmdutils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}

	users := []*models.UserForListDto{}
	for {
		response, err := apiClient.Client.Users.UsersList(params, apiClient)
		if err != nil {
			return err
		}
		users = append(users, response.Payload.Data...)
		usersCount := int32(len(users))
		if opts.Limit != 0 && usersCount >= opts.Limit {
			break
		}
		if usersCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&usersCount)
	}

	if opts.Limit != 0 && int32(len(users)) > opts.Limit {
		users = users[:opts.Limit]
	}

	cmdutils.PrettyPrint(users)
	return
}
