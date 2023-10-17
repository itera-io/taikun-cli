package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

// ListFields defines a slice of fields corresponding to the columns in the output.
// Some columns are set as visible by default and some are hidden by default.
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

// NewCmdList creates and returns a cobra command for listing users.
// It supports Sorting (with Reverse) and Limiting the output
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

// listRun calls the API, gets the Users and prints them in a table.
func listRun(opts *ListOptions) (err error) {
	users, err := ListUsers(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(users, ListFields)
}

// ListUsers sends multiple queries to the API and returns a list of users.
// Users are returned in the UserForListDto structs generated in models.
func ListUsers(opts *ListOptions) (userList []taikuncore.UserForListDto, err error) {
	// Prepare the request
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.UsersAPI.UsersList(context.TODO())
	// Set Organization ID if it is set in command line options
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}
	// Set Sorting if set in command line options
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	// Initialise a new, empty slice of UserForListDto structs generated in models.
	userList = make([]taikuncore.UserForListDto, 0)

	// Execute the request, it returns 50 users and then execute it again with an Offset until you have read all of it.
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return nil, tk.CreateError(response, err)
		}

		userList = append(userList, data.GetData()...)
		usersCount := int32(len(userList))

		// We have (over)reached the limit, cut it at the limit and break
		if opts.Limit != 0 && usersCount >= opts.Limit {
			if int32(len(userList)) > opts.Limit {
				userList = userList[:opts.Limit]
			}
			break
		}
		// We have read all the users
		if usersCount == data.GetTotalCount() {
			break
		}

		// The new request will be shifted to the next page of users.
		myRequest = myRequest.Offset(usersCount)
	}

	return userList, nil
}
