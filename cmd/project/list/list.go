package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
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
		field.NewVisible(
			"STATUS", "status",
		),
		field.NewVisible(
			"HEALTH", "health",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewVisible(
			"CLOUD", "cloudType",
		),
		field.NewVisible(
			"QUOTA-ID", "quotaId",
		),
		field.NewVisible(
			"EXPIRES", "expiredAt",
		),
		field.NewVisible(
			"LOCK", "isLocked",
		),
	},
	// TODO FORMAT???
	// TODO check JSON
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List projects",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByAndReverseFlags(&cmd, "projects", listFields)
	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := projects.NewProjectsListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var projects = make([]*models.ProjectListForUIDto, 0)
	for {
		response, err := apiClient.Client.Projects.ProjectsList(params, apiClient)
		if err != nil {
			return err
		}
		projects = append(projects, response.Payload.Data...)
		projectsCount := int32(len(projects))
		if config.Limit != 0 && projectsCount >= config.Limit {
			break
		}
		if projectsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&projectsCount)
	}

	if config.Limit != 0 && int32(len(projects)) > config.Limit {
		projects = projects[:config.Limit]
	}

	out.PrintResults(projects, listFields)
	return
}
