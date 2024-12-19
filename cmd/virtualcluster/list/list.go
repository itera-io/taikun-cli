package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
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
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ACCESS-IP", "accessIp",
		),
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewHidden(
			"HEALTH", "health",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
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
		Use:   "list [project-id]",
		Short: "List virtual clusters for a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(args[0], &opts) // List user by ID
		},
		Args:    cobra.ExactArgs(1),
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(cmd, "virtualcluster", ListFields)
	cmdutils.AddColumnsFlag(cmd, ListFields)
	cmdutils.AddLimitFlag(cmd, &opts.Limit)

	return cmd
}

// listRun calls the API, gets the Users and prints them in a table.
func listRun(projectIdString string, opts *ListOptions) (err error) {
	projectId, err := types.Atoi32(projectIdString)
	if err != nil {
		return err
	}

	// Prepare the request
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.VirtualClusterAPI.VirtualClusterList(context.TODO(), projectId)

	// Set Sorting if set in command line options
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	// Initialise a new, empty slice of VClusterListDto structs generated in models.
	vcsList := make([]taikuncore.VClusterListDto, 0)

	// Execute the request, it returns 50 vcs and then execute it again with an Offset until you have read all of it.
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		vcsList = append(vcsList, data.GetData()...)
		vcsCount := int32(len(vcsList))

		// We have (over)reached the limit, cut it at the limit and break
		if opts.Limit != 0 && vcsCount >= opts.Limit {
			if int32(len(vcsList)) > opts.Limit {
				vcsList = vcsList[:opts.Limit]
			}
			break
		}
		// We have read all the vcs
		if vcsCount == data.GetTotalCount() {
			break
		}

		// The new request will be shifted to the next page of vcs.
		myRequest = myRequest.Offset(vcsCount)
	}

	return out.PrintResults(vcsList, ListFields)
}
