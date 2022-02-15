package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"LABEL", "label",
		),
		field.NewVisible(
			"VALUE", "value",
		),
	},
)

type ListOptions struct {
	ShowbackRuleID int32
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list <showback-rule-id>",
		Short: "list a showback rule's labels",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ShowbackRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	showbackRule, err := GetShowbackRuleByID(opts.ShowbackRuleID)
	if err != nil {
		return
	}

	labels := showbackRule.Labels

	if opts.Limit != 0 && int32(len(labels)) > opts.Limit {
		labels = labels[:opts.Limit]
	}

	return out.PrintResults(labels, listFields)
}

func GetShowbackRuleByID(showbackRuleID int32) (showbackRule *models.ShowbackRulesListDto, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := showback.NewShowbackRulesListParams().WithV(api.Version)
	params = params.WithID(&showbackRuleID)

	response, err := apiClient.Client.Showback.ShowbackRulesList(params, apiClient)
	if err != nil {
		return
	}
	if len(response.Payload.Data) != 1 {
		return nil, cmderr.ResourceNotFoundError("Showback rule", showbackRuleID)
	}

	showbackRule = response.Payload.Data[0]
	return
}
