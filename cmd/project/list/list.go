package list

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/config"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/projects"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
}

func printResults(projects []*models.ProjectListForUIDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(projects)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(projects))
		for i, project := range projects {
			data[i] = project
		}
		format.PrettyPrintTable(data,
			"id",
			"name",
			"organizationName",
			"status",
			"health",
			"createdAt",
			"kubernetesCurrentVersion",
			"cloudType",
			"hasKubeConfigFile",
			"quotaId",
			"expiredAt",
			"isLocked",
		)
	}
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByFlag(cmd, &opts.SortBy, models.ProjectListForUIDto{})

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := projects.NewProjectsListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var projects = make([]*models.ProjectListForUIDto, 0)
	for {
		response, err := apiClient.Client.Projects.ProjectsList(params, apiClient)
		if err != nil {
			return err
		}
		projects = append(projects, response.Payload.Data...)
		projectsCount := int32(len(projects))
		if opts.Limit != 0 && projectsCount >= opts.Limit {
			break
		}
		if projectsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&projectsCount)
	}

	if opts.Limit != 0 && int32(len(projects)) > opts.Limit {
		projects = projects[:opts.Limit]
	}

	printResults(projects)
	return
}
