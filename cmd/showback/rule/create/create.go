package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	GlobalAlertLimit     int32
	Kind                 string
	MetricName           string
	Name                 string
	OrganizationID       int32
	Price                float64
	ProjectAlertLimit    int32
	Type                 string
	ShowbackCredentialID int32
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <name>",
		Short: "Create a showback rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if err := cmdutils.CheckFlagValue("kind", opts.Kind, types.ShowbackKinds); err != nil {
				return err
			}
			if err := cmdutils.CheckFlagValue("type", opts.Type, types.PrometheusTypes); err != nil {
				return err
			}
			return createRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.GlobalAlertLimit, "global-alert-limit", "g", 0, "Global alert limit")
	cmdutils.MarkFlagRequired(&cmd, "global-alert-limit")

	cmd.Flags().StringVarP(&opts.Kind, "kind", "k", "", "Kind (required)")
	cmdutils.MarkFlagRequired(&cmd, "kind")
	cmdutils.RegisterStaticFlagCompletion(&cmd, "kind", types.ShowbackKinds.Keys()...)

	cmd.Flags().StringVarP(&opts.MetricName, "metric-name", "m", "", "Metric name (required)")
	cmdutils.MarkFlagRequired(&cmd, "metric-name")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmd.Flags().Float64VarP(&opts.Price, "price", "p", 0, "Price")
	cmdutils.MarkFlagRequired(&cmd, "price")

	cmd.Flags().StringVarP(&opts.Type, "type", "t", "", "Type (required)")
	cmdutils.MarkFlagRequired(&cmd, "type")
	cmdutils.RegisterStaticFlagCompletion(&cmd, "type", types.PrometheusTypes.Keys()...)

	cmd.Flags().Int32Var(&opts.ProjectAlertLimit, "project-alert-limit", 0, "Project alert limit")

	cmd.Flags().Int32VarP(&opts.ShowbackCredentialID, "showback-credential-id", "s", 0, "Showback credential ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CreateShowbackRuleCommand{
		GlobalAlertLimit: opts.GlobalAlertLimit,
		Kind:             types.GetShowbackKind(opts.Kind),
		MetricName:       opts.MetricName,
		Name:             opts.Name,
		Price:            opts.Price,
		Type:             types.GetPrometheusType(opts.Type),
	}

	if opts.OrganizationID != 0 {
		body.OrganizationID = opts.OrganizationID
	}

	if opts.ProjectAlertLimit != 0 {
		body.ProjectAlertLimit = opts.ProjectAlertLimit
	}

	if opts.ShowbackCredentialID != 0 {
		body.ShowbackCredentialID = &opts.ShowbackCredentialID
	}

	params := showback.NewShowbackCreateRuleParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Showback.ShowbackCreateRule(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"metricName",
			"organizationName",
			"kind",
			"type",
			"globalAlertLimit",
			"projectAlertLimit",
			"price",
			"showbackCredentialName",
			"createdAt",
		)
	}

	return
}
