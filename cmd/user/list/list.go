package list

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var ListFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewHidden(
			"DISPLAY-NAME", "displayName",
		),
		field.NewVisible(
			"EMAIL", "email",
		),
		field.NewVisible(
			"ROLE", "globalRole",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"2FA", "is2FAEnabled",
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
			return listRun(cmd, &opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddOrgIDFlag(cmd, &opts.OrganizationID)
	cmdutils.AddColumnsFlag(cmd, ListFields)
	cmdutils.AddLimitFlag(cmd, &opts.Limit)

	return cmd
}

func listRun(cmd *cobra.Command, opts *ListOptions) (err error) {
	orgID, err := cmdutils.ResolveOrgID(opts.OrganizationID, cmdutils.IsRobotAuth())
	if err != nil {
		return err
	}
	opts.OrganizationID = orgID
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()
	users, err := ListUsers(ctx, opts)
	if err != nil {
		return err
	}

	return out.PrintResults(users, ListFields)
}

func ListUsers(ctx context.Context, opts *ListOptions) ([]interface{}, error) {
	myApiClient := tk.NewClient()

	// Get current user's account ID for detail fetches
	userInfo, response, err := myApiClient.Client.UsersAPI.UsersUserInfo(ctx).Execute()
	if err != nil {
		return nil, tk.CreateError(response, err)
	}
	accountID := userInfo.Data.Account.AccountId

	// Fetch user IDs via dropdown (paginated)
	myRequest := myApiClient.Client.UsersAPI.UsersDropdown(ctx)
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	var dropdownUsers []taikuncore.CommonStringBasedDropdownDto
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return nil, tk.CreateError(response, err)
		}

		dropdownUsers = append(dropdownUsers, data.GetData()...)
		count := int32(len(dropdownUsers))

		if opts.Limit != 0 && count >= opts.Limit {
			if count > opts.Limit {
				dropdownUsers = dropdownUsers[:opts.Limit]
			}
			break
		}
		if count == int32(data.GetTotalCount()) {
			break
		}
		myRequest = myRequest.Offset(count)
	}

	// Fetch full details for each user
	results := make([]interface{}, 0, len(dropdownUsers))
	for _, u := range dropdownUsers {
		userID := u.GetId()
		detail, response, err := myApiClient.Client.AccountsAPI.AccountsAccountUserDetails(ctx, accountID, userID).Execute()
		if err != nil {
			_ = tk.CreateError(response, err)
			continue
		}
		results = append(results, detail)
	}

	return results, nil
}
