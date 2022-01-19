package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"METRIC", "metricName",
		),
		field.NewVisible(
			"PRICE", "price",
		),
		field.NewVisible(
			"KIND", "kind",
		),
		field.NewVisible(
			"TYPE", "type",
		),
		field.NewVisible(
			"GLOBAL-ALERT-LIMIT", "globalAlertLimit",
		),
		field.NewVisible(
			"PROJECT-ALERT-LIMIT", "projectAlertLimit",
		),
		field.NewVisible(
			"CREDENTIAL", "showbackCredentialName",
		),
		field.NewVisible(
			"CREDENTIAL-ID", "showbackCredentialId",
		),
		field.NewVisibleWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
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

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a showback rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if err := cmdutils.CheckFlagValue("kind", opts.Kind, types.ShowbackKinds); err != nil {
				return err
			}
			if err := cmdutils.CheckFlagValue("type", opts.Type, types.PrometheusTypes); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.GlobalAlertLimit, "global-alert-limit", "g", 0, "Global alert limit")
	cmdutils.MarkFlagRequired(&cmd, "global-alert-limit")

	cmd.Flags().StringVarP(&opts.Kind, "kind", "k", "", "Kind (required)")
	cmdutils.MarkFlagRequired(&cmd, "kind")
	cmdutils.SetFlagCompletionValues(&cmd, "kind", types.ShowbackKinds.Keys()...)

	cmd.Flags().StringVarP(&opts.MetricName, "metric-name", "m", "", "Metric name (required)")
	cmdutils.MarkFlagRequired(&cmd, "metric-name")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmd.Flags().Float64VarP(&opts.Price, "price", "p", 0, "Price")
	cmdutils.MarkFlagRequired(&cmd, "price")

	cmd.Flags().StringVarP(&opts.Type, "type", "t", "", "Type (required)")
	cmdutils.MarkFlagRequired(&cmd, "type")
	cmdutils.SetFlagCompletionValues(&cmd, "type", types.PrometheusTypes.Keys()...)

	cmd.Flags().Int32Var(&opts.ProjectAlertLimit, "project-alert-limit", 0, "Project alert limit")

	cmd.Flags().Int32VarP(&opts.ShowbackCredentialID, "showback-credential-id", "s", 0, "Showback credential ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
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

	params := showback.NewShowbackCreateRuleParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Showback.ShowbackCreateRule(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload, addFields)
	}

	return
}
