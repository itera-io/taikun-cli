package bind

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/spf13/cobra"
)

type BindOptions struct {
	BillingRuleID   int32
	OrganizationIDs []int32
}

func NewCmdBind() *cobra.Command {
	var opts BindOptions

	cmd := cobra.Command{
		Use:   "bind <organization-id>...",
		Short: "Bind a billing rule to one or more organizations",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ids, err := cmdutils.ArgsToNumericalIDs(args)
			if err != nil {
				return err
			}
			opts.OrganizationIDs = ids
			return bindRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.BillingRuleID, "billing-rule-id", "b", 0, "Billing rule ID")
	cmdutils.MarkFlagRequired(&cmd, "billing-rule-id")

	return &cmd
}

func bindRun(opts *BindOptions) (err error) {
	// FIXME
	return
}
