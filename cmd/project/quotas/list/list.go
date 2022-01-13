package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/project_quotas"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects quotas",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByAndReverseFlags(cmd, models.KubernetesProfilesListDto{})

	return cmd
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
	if config.ReverseSortDirection {
		api.ReverseSortDirection()
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(&api.SortDirection)
	}

	var projectQuotas = make([]*models.ProjectQuotaListDto, 0)
	for {
		response, err := apiClient.Client.ProjectQuotas.ProjectQuotasList(params, apiClient)
		if err != nil {
			return err
		}
		projectQuotas = append(projectQuotas, response.Payload.Data...)
		count := int32(len(projectQuotas))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(projectQuotas)) > config.Limit {
		projectQuotas = projectQuotas[:config.Limit]
	}

	out.PrintResults(projectQuotas,
		"id",
		"cpu",
		"isCpuUnlimited",
		"diskSize",
		"isDiskSizeUnlimited",
		"ram",
		"isRamUnlimited",
	)
	return
}
