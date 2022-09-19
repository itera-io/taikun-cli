package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/users"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var ListFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "username",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"DISPLAY-NAME", "displayName",
		),
		field.NewVisible(
			"ROLE", "role",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewHidden(
			"OWNER", "owner",
		),
		field.NewHidden(
			"EMAIL", "email",
		),
		field.NewHidden(
			"EMAIL-CONFIRMED", "isEmailConfirmed",
		),
		field.NewHidden(
			"EMAIL-NOTIFICATIONS", "isEmailNotificationEnabled",
		),
		field.NewHidden(
			"APPROVED-BY-PARTNER", "isApprovedByPartner",
		),
		field.NewHidden(
			"CSM", "isCsm",
		),
		field.NewHidden(
			"SUBSCRIPTION-UPDATES", "isEligibleUpdateSubscription",
		),
		field.NewHidden(
			"MUST-RESET-PASSWORD", "isForcedToResetPassword",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"READ-ONLY", "isReadOnly",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED", "createdAt", out.FormatDateTimeString,
		),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
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

	cmdutils.AddSortByAndReverseFlags(cmd, "users", ListFields)
	cmdutils.AddColumnsFlag(cmd, ListFields)
	cmdutils.AddLimitFlag(cmd, &opts.Limit)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	users, err := ListUsers(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(users, ListFields)
}

func ListUsers(opts *ListOptions) (userList []*models.UserForListDto, err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return nil, err
	}

	params := users.NewUsersListParams().WithV(taikungoclient.Version)
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
		if opts.Limit != 0 && usersCount >= opts.Limit {
			break
		}

		if usersCount == response.Payload.TotalCount {
			break
		}

		params = params.WithOffset(&usersCount)
	}

	if opts.Limit != 0 && int32(len(userList)) > opts.Limit {
		userList = userList[:opts.Limit]
	}

	return userList, nil
}
