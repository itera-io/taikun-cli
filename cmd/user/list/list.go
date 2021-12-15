package list

import (
	"taikun-cli/api"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return utils.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

	return cmd
}

func printResults(users []*models.UserForListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		utils.PrettyPrintJson(users)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(users))
		for i, user := range users {
			data[i] = user
		}
		utils.PrettyPrintTable(data,
			"id",
			"username",
			"role",
			"organizationName",
			"email",
			"isEmailConfirmed",
			"isEmailNotificationEnabled",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := users.NewUsersListParams().WithV(utils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		utils.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&utils.SortDirection)
	}

	var users = make([]*models.UserForListDto, 0)
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

	printResults(users)
	return
}
