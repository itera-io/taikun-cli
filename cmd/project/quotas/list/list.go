package list

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/config"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/project_quotas"
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
		Short: "List project quotas",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return cmderr.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return cmderr.OutputFormatInvalidError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByFlag(cmd, &opts.SortBy, models.KubernetesProfilesListDto{})

	return cmd
}

func printResults(projectQuotas []*models.ProjectQuotaListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(projectQuotas)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(projectQuotas))
		for i, projectQuota := range projectQuotas {
			data[i] = projectQuota
		}
		format.PrettyPrintTable(data,
			"id",
			"cpu",
			"isCpuUnlimited",
			"diskSize",
			"isDiskSizeUnlimited",
			"ram",
			"isRamUnlimited",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := project_quotas.NewProjectQuotasListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
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

	printResults(projectQuotas)
	return
}
