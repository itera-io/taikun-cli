package add

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient"
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
		Use:   "add <showback-rule-id>",
		Short: "Add a label to a showback rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.ShowbackRuleID, err = types.Atoi32(args[0])
			if err != nil {
				return cmderr.ErrIDArgumentNotANumber
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

func addRun(opts *AddOptions) error {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return err
	}

	showbackRule, err := list.GetShowbackRuleByID(opts.ShowbackRuleID)
	if err != nil {
		return err
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

	params := showback.NewShowbackUpdateRuleParams().WithV(taikungoclient.Version)
	params = params.WithBody(&body)

	if _, err := apiClient.Client.Showback.ShowbackUpdateRule(params, apiClient); err != nil {
		return err
	}

	out.PrintStandardSuccess()

	return nil
}
