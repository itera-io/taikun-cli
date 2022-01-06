package clear

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ClearOptions struct {
	ShowbackRuleID int32
}

func NewCmdClear() *cobra.Command {
	var opts ClearOptions

	cmd := cobra.Command{
		Use:   "clear",
		Short: "clear a showback rule's labels",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ShowbackRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return clearRun(&opts)
		},
	}

	return &cmd
}

func clearRun(opts *ClearOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	showbackRule, err := list.GetShowbackRuleByID(opts.ShowbackRuleID)
	if err != nil {
		return
	}
	emptyLabels := make([]*models.ShowbackLabelCreateDto, 0)

	body := models.UpdateShowbackRuleCommand{
		GlobalAlertLimit:  showbackRule.GlobalAlertLimit,
		ID:                opts.ShowbackRuleID,
		Kind:              types.GetShowbackKind(showbackRule.Kind),
		Labels:            emptyLabels,
		MetricName:        showbackRule.MetricName,
		Name:              showbackRule.Name,
		Price:             showbackRule.Price,
		ProjectAlertLimit: showbackRule.ProjectAlertLimit,
		Type:              types.GetPrometheusType(showbackRule.Type),
	}

	params := showback.NewShowbackUpdateRuleParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Showback.ShowbackUpdateRule(params, apiClient)
	if err == nil {
		format.PrintStandardSuccess()
	}

	return
}
