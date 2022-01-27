package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/project_quotas"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"PROJECT", "projectName",
		),
		field.NewHidden(
			"PROJECT-ID", "projectId",
		),
		field.NewVisibleWithToStringFunc(
			"CPU", "cpu", out.FormatNumber,
		),
		field.NewVisible(
			"UNLIMITED-CPU", "isCpuUnlimited",
		),
		field.NewVisibleWithToStringFunc(
			"DISK", "diskSize", out.FormatBToGiB,
		),
		field.NewVisible(
			"UNLIMITED-DISK", "isDiskSizeUnlimited",
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "ram", out.FormatBToGiB,
		),
		field.NewVisible(
			"UNLIMITED-RAM", "isRamUnlimited",
		),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List project quotas",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "project-quotas", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := project_quotas.NewProjectQuotasListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var projectQuotas = make([]*models.ProjectQuotaListDto, 0)
	for {
		response, err := apiClient.Client.ProjectQuotas.ProjectQuotasList(params, apiClient)
		if err != nil {
			return err
		}
		projectQuotas = append(projectQuotas, response.Payload.Data...)
		count := int32(len(projectQuotas))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(projectQuotas)) > opts.Limit {
		projectQuotas = projectQuotas[:opts.Limit]
	}

	return out.PrintResults(projectQuotas, listFields)
}
