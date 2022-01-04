package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/list"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	ShowbackRuleID int32
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
	}

	cmdutils.AddLimitFlag(&cmd)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	labels, err := GetShowbackRuleLabelsByID(opts.ShowbackRuleID)
	if err != nil {
		return
	}

	if list.Limit != 0 && int32(len(labels)) > list.Limit {
		labels = labels[:list.Limit]
	}

	format.PrintResults(labels, "label", "value")

	return
}

func GetShowbackRuleLabelsByID(id int32) (labels []*models.ShowbackLabelCreateDto, err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := showback.NewShowbackRulesListParams().WithV(apiconfig.Version)
	params = params.WithID(&id)

	response, err := apiClient.Client.Showback.ShowbackRulesList(params, apiClient)
	if err != nil {
		return
	}
	if len(response.Payload.Data) != 1 {
		return nil, cmderr.ResourceNotFoundError("Showback rule", id)
	}

	labels = response.Payload.Data[0].Labels
	return
}
