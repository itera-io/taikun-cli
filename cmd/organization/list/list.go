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

var ListFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"FULL-NAME", "fullName",
		),
		field.NewHidden(
			"EMAIL", "email",
		),
		field.NewHidden(
			"CLOUD-CREDENTIALS", "cloudCredentials",
		),
		field.NewHidden(
			"PROJECTS", "projects",
		),
		field.NewHidden(
			"SERVERS", "servers",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
	},
)

type ListOptions struct {
	Limit int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List organizations",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(cmd, &opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddSortByAndReverseFlags(&cmd, "organizations", ListFields)
	cmdutils.AddColumnsFlag(&cmd, ListFields)
	cmdutils.AddLimitFlag(&cmd, &opts.Limit)

	return &cmd
}

func listRun(cmd *cobra.Command, opts *ListOptions) (err error) {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	organizations, err := ListOrganizations(ctx, opts)
	if err != nil {
		return err
	}

	return out.PrintResults(organizations, ListFields)
}

func ListOrganizations(ctx context.Context, opts *ListOptions) (organizations []taikuncore.OrganizationDetailsDto, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.OrganizationsAPI.OrganizationsList(ctx)
	// Set Sorting if set in command line options
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	// Execute the request, it returns only 50 lines in one page
	// then execute it again with an Offset until you have read all of it.
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return nil, tk.CreateError(response, err)
		}
		organizations = append(organizations, data.Data...)

		organizationsCount := int32(len(organizations))
		if opts.Limit != 0 && organizationsCount >= opts.Limit {
			break
		}

		if int64(organizationsCount) == data.GetTotalCount() { // casting is safe, extending maximum bounds
			break
		}

		myRequest = myRequest.Offset(organizationsCount)
	}

	// We have (over)reached the limit, cut it at the limit and break
	if opts.Limit != 0 && int32(len(organizations)) > opts.Limit {
		organizations = organizations[:opts.Limit]
	}

	return organizations, nil
}
