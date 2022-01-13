package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	ShowbackRuleID int32
	Label          string
	Value          string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add",
		Short: "Add a label to a showback rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ShowbackRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.IDArgumentNotANumberError
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Label, "label", "l", "", "Label")
	cmdutils.MarkFlagRequired(&cmd, "label")

	cmd.Flags().StringVarP(&opts.Value, "value", "v", "", "Value")
	cmdutils.MarkFlagRequired(&cmd, "value")

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	showbackRule, err := list.GetShowbackRuleByID(opts.ShowbackRuleID)
	if err != nil {
		return
	}
	newLabel := models.ShowbackLabelCreateDto{Label: opts.Label, Value: opts.Value}
	showbackRule.Labels = append(showbackRule.Labels, &newLabel)

	body := models.UpdateShowbackRuleCommand{
		GlobalAlertLimit:  showbackRule.GlobalAlertLimit,
		ID:                opts.ShowbackRuleID,
		Kind:              types.GetShowbackKind(showbackRule.Kind),
		Labels:            showbackRule.Labels,
		MetricName:        showbackRule.MetricName,
		Name:              showbackRule.Name,
		Price:             showbackRule.Price,
		ProjectAlertLimit: showbackRule.ProjectAlertLimit,
		Type:              types.GetPrometheusType(showbackRule.Type),
	}

	params := showback.NewShowbackUpdateRuleParams().WithV(api.Version)
	params = params.WithBody(&body)

	_, err = apiClient.Client.Showback.ShowbackUpdateRule(params, apiClient)
	if err == nil {
		out.PrintStandardSuccess()
	}

	return
}
