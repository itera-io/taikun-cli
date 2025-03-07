package clear

import (
	"context"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	showbackRule, err := list.GetShowbackRuleByID(opts.ShowbackRuleID)
	if err != nil {
		return err
	}
	emptyLabels := make([]taikunshowback.ShowbackLabelCreateDto, 0)
	body := taikunshowback.UpdateShowbackRuleCommand{
		Id:                &opts.ShowbackRuleID,
		Name:              showbackRule.Name,
		MetricName:        showbackRule.MetricName,
		Kind:              types.GetShowbackKind(showbackRule.GetKind()),
		Type:              types.GetEPrometheusType(showbackRule.GetType()),
		Price:             *taikunshowback.NewNullableFloat64(showbackRule.Price),
		ProjectAlertLimit: *taikunshowback.NewNullableInt32(showbackRule.ProjectAlertLimit),
		Labels:            emptyLabels,
	}

	// Execute a query into the API + graceful exit
	_, response, err := myApiClient.ShowbackClient.ShowbackRulesAPI.ShowbackrulesUpdate(context.TODO()).UpdateShowbackRuleCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil

}
