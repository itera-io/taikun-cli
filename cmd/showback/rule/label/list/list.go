package list

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
				return cmderr.ErrIDArgumentNotANumber
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

func GetShowbackRuleByID(showbackRuleID int32) (showbackRule *taikunshowback.ShowbackRulesListDto, err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.ShowbackClient.ShowbackRulesAPI.ShowbackrulesList(context.TODO()).Id(showbackRuleID).Execute()
	if err != nil {
		return nil, tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	if len(data.GetData()) != 1 {
		return nil, cmderr.ResourceNotFoundError("Showback rule", showbackRuleID)
	}

	showbackRule = &data.Data[0]
	return showbackRule, nil
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := showback_rules.NewShowbackRulesListParams().WithV(taikungoclient.Version)
		params = params.WithID(&showbackRuleID)

		response, err := apiClient.ShowbackClient.ShowbackRules.ShowbackRulesList(params, apiClient)
		if err != nil {
			return
		}

		if len(response.Payload.Data) != 1 {
			return nil, cmderr.ResourceNotFoundError("Showback rule", showbackRuleID)
		}

		showbackRule = response.Payload.Data[0]

		return
	*/
}
