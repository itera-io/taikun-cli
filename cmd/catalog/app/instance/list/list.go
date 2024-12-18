package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

// ListFields defines a slice of fields corresponding to the columns in the output.
// Some columns are set as visible by default and some are hidden by default.
var ListFields = fields.New(
	[]*field.Field{
		field.NewVisible("ID", "id"),
		field.NewVisible("NAME", "name"),
		field.NewVisible("APP-NAME", "catalogAppName"),
		field.NewVisible("APP-REPO", "appRepoName"),
		field.NewVisible("NAMESPACE", "namespace"),
		field.NewVisible("STATUS", "status"),
		field.NewVisible("VERSION", "version"),
		field.NewHidden("CATALOG-NAME", "catalogName"),
		field.NewHidden("CATALOG-ID", "catalogId"),
		field.NewHiddenWithToStringFunc("CREATED", "created", out.FormatDateTimeString),
	},
)

type ListOptions struct {
	projectId      int32
	OrganizationID int32
	Limit          int32
}

// NewCmdList creates and returns a cobra command for listing users.
// It supports Sorting (with Reverse) and Limiting the output
func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list <PROJECT_ID>",
		Short: "List application instances for project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.projectId, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return listRun(&opts)
		},
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
	users, err := ListAppInstances(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(users, ListFields)
}

func ListAppInstances(opts *ListOptions) (appInstanceList []taikuncore.InstanceAppListDto, err error) {
	// Prepare the request
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ProjectAppsAPI.ProjectappList(context.TODO()).ProjectId(opts.projectId)
	// Set Organization ID if it is set in command line options
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}
	// Set Sorting if set in command line options
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	appInstanceList = make([]taikuncore.InstanceAppListDto, 0)

	// Execute the request, it returns 50 users and then execute it again with an Offset until you have read all of it.
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return nil, tk.CreateError(response, err)
		}

		appInstanceList = append(appInstanceList, data.GetData()...)
		appInstancesCount := int32(len(appInstanceList))

		// We have (over)reached the limit, cut it at the limit and break
		if opts.Limit != 0 && appInstancesCount >= opts.Limit {
			if int32(len(appInstanceList)) > opts.Limit {
				appInstanceList = appInstanceList[:opts.Limit]
			}
			break
		}
		// We have read all the users
		if appInstancesCount == data.GetTotalCount() {
			break
		}

		// The new request will be shifted to the next page of users.
		myRequest = myRequest.Offset(appInstancesCount)
	}

	return appInstanceList, nil
}
