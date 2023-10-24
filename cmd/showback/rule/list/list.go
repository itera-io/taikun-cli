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
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"METRIC", "metricName",
		),
		field.NewVisible(
			"PRICE", "price",
		),
		field.NewVisible(
			"KIND", "kind",
		),
		field.NewVisible(
			"TYPE", "type",
		),
		field.NewVisible(
			"GLOBAL-ALERT-LIMIT", "globalAlertLimit",
		),
		field.NewVisible(
			"PROJECT-ALERT-LIMIT", "projectAlertLimit",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
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
		Short: "List showback rules",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "showback-rules", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	myRequest := myApiClient.ShowbackClient.ShowbackRulesAPI.ShowbackrulesList(context.TODO())
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	var showbackRules = make([]taikunshowback.ShowbackRulesListDto, 0)

	// Send the query and move offset until you get all the data
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		showbackRules = append(showbackRules, data.GetData()...)

		count := int32(len(showbackRules))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(showbackRules)) > opts.Limit {
		showbackRules = showbackRules[:opts.Limit]
	}

	return out.PrintResults(showbackRules, listFields)
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := showback_rules.NewShowbackRulesListParams().WithV(taikungoclient.Version)
		if opts.OrganizationID != 0 {
			params = params.WithOrganizationID(&opts.OrganizationID)
		}

		if config.SortBy != "" {
			params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
		}

		var showbackRules = make([]*models.ShowbackRulesListDto, 0)

		for {
			response, err := apiClient.ShowbackClient.ShowbackRules.ShowbackRulesList(params, apiClient)
			if err != nil {
				return err
			}

			showbackRules = append(showbackRules, response.Payload.Data...)

			count := int32(len(showbackRules))
			if opts.Limit != 0 && count >= opts.Limit {
				break
			}

			if count == response.Payload.TotalCount {
				break
			}

			params = params.WithOffset(&count)
		}

		if opts.Limit != 0 && int32(len(showbackRules)) > opts.Limit {
			showbackRules = showbackRules[:opts.Limit]
		}

		return out.PrintResults(showbackRules, listFields)
	*/
}
