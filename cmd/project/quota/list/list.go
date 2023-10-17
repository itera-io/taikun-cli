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

var listFields = fields.New(
	[]*field.Field{
		//field.NewVisible(
		//	"ID", "id",
		//),
		field.NewVisible(
			"PROJECT-ID", "projectId",
		),
		field.NewVisible(
			"PROJECT", "projectName",
		),
		field.NewVisibleWithToStringFunc(
			"CPU", "serverCpu", out.FormatNumberInteger,
		),
		field.NewVisibleWithToStringFunc(
			"RAM", "serverRam", out.FormatBToGiB,
		),
		field.NewVisibleWithToStringFunc(
			"DISK", "serverDiskSize", out.FormatBToGiB,
		),
		field.NewVisibleWithToStringFunc(
			"VM-CPU", "vmCpu", out.FormatNumberInteger,
		),
		field.NewVisibleWithToStringFunc(
			"VM-RAM", "vmRam", out.FormatBToGiB,
		),
		field.NewVisibleWithToStringFunc(
			"VM-VOLUME-SIZE", "vmVolumeSize", out.FormatNumberAddGibString, // No conversion needed. API takes GBs
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
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.ProjectQuotasAPI.ProjectquotasList(context.TODO())
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	projectQuotas := make([]taikuncore.ProjectQuotaListDto, 0)

	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		projectQuotas = append(projectQuotas, data.GetData()...)

		count := int32(len(projectQuotas))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(projectQuotas)) > opts.Limit {
		projectQuotas = projectQuotas[:opts.Limit]
	}

	return out.PrintResults(projectQuotas, listFields)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := project_quotas.NewProjectQuotasListParams().WithV(taikungoclient.Version)
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
	*/
}
