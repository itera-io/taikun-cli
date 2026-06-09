package add

import (
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/showback/rule/label/list"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	tk "github.com/itera-io/taikungoclient"
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
			return addRun(cmd, &opts)
		},
	}

	cmd.Flags().StringVarP(&opts.Label, "label", "l", "", "Label")
	cmdutils.MarkFlagRequired(&cmd, "label")

	cmd.Flags().StringVarP(&opts.Value, "value", "v", "", "Value")
	cmdutils.MarkFlagRequired(&cmd, "value")

	return &cmd
}

func addRun(cmd *cobra.Command, opts *AddOptions) error {
	ctx, cancel := cmdutils.APIContext(cmd)
	defer cancel()

	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	showbackRule, err := list.GetShowbackRuleByID(cmd, opts.ShowbackRuleID)
	if err != nil {
		return err
	}
	newLabel := taikunshowback.ShowbackLabelCreateDto{
		Label: *taikunshowback.NewNullableString(&opts.Label),
		Value: *taikunshowback.NewNullableString(&opts.Value),
	}
	showbackRule.Labels = append(showbackRule.Labels, newLabel)
	body := taikunshowback.UpdateShowbackRuleCommand{
		Id:                &opts.ShowbackRuleID,
		Name:              showbackRule.Name,
		MetricName:        showbackRule.MetricName,
		Kind:              types.GetShowbackKind(showbackRule.GetKind()),
		Type:              types.GetEPrometheusType(showbackRule.GetType()),
		Price:             *taikunshowback.NewNullableFloat64(showbackRule.Price),
		ProjectAlertLimit: *taikunshowback.NewNullableInt32(showbackRule.ProjectAlertLimit),
		GlobalAlertLimit:  *taikunshowback.NewNullableInt32(showbackRule.GlobalAlertLimit),
		Labels:            showbackRule.Labels,
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.ShowbackClient.ShowbackRulesAPI.ShowbackrulesUpdate(ctx).UpdateShowbackRuleCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	out.PrintStandardSuccess()
	return nil

}
