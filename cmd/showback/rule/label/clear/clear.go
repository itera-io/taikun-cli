package clear

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
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
		Use:   "clear <showback-rule-id>",
		Short: "clear a showback rule's labels",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ShowbackRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
			}
			return clearRun(&opts)
		},
	}

	return &cmd
}

func clearRun(opts *ClearOptions) error {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return err
	}

	showbackRule, err := list.GetShowbackRuleByID(opts.ShowbackRuleID)
	if err != nil {
		return err
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

	params := showback.NewShowbackUpdateRuleParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	if _, err := apiClient.Client.Showback.ShowbackUpdateRule(params, apiClient); err != nil {
		return err
	}

	out.PrintStandardSuccess()

	return nil
}
