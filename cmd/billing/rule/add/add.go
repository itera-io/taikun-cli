package add

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/prometheus"
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
			"METRIC", "metricName",
		),
		field.NewVisible(
			"PRICE", "price",
		),
		field.NewVisible(
			"TYPE", "type",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	BillingCredentialID int32
	Labels              []string
	MetricName          string
	Name                string
	Price               float64
	PriceRate           int32
	Type                string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a billing rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if err := cmdutils.CheckFlagValue("type", opts.Type, types.PrometheusTypes); err != nil {
				return err
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().Int32VarP(&opts.BillingCredentialID, "billing-credential-id", "b", 0, "Billing credential ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "billing-credential-id")

	cmd.Flags().StringSliceVarP(&opts.Labels, "labels", "l", []string{}, "Labels (format: \"label=value,label2=value2,...\")")
	cmdutils.MarkFlagRequired(&cmd, "labels")

	cmd.Flags().StringVarP(&opts.MetricName, "metric-name", "m", "", "Metric name (required)")
	cmdutils.MarkFlagRequired(&cmd, "metric-name")

	cmd.Flags().Float64Var(&opts.Price, "price", 0, "Price (required)")
	cmdutils.MarkFlagRequired(&cmd, "price")

	cmd.Flags().Int32Var(&opts.PriceRate, "price-rate", 0, "Price rate (required)")
	cmdutils.MarkFlagRequired(&cmd, "price-rate")

	cmd.Flags().StringVarP(&opts.Type, "type", "t", "", "Type (required)")
	cmdutils.MarkFlagRequired(&cmd, "type")
	cmdutils.SetFlagCompletionValues(&cmd, "type", types.PrometheusTypes.Keys()...)

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.RuleCreateCommand{
		MetricName:            opts.MetricName,
		Name:                  opts.Name,
		OperationCredentialID: opts.BillingCredentialID,
		Price:                 opts.Price,
		RuleDiscountRate:      opts.PriceRate,
		Type:                  types.GetPrometheusType(opts.Type),
	}

	body.Labels, err = parseLabelsFlag(opts.Labels)
	if err != nil {
		return
	}

	params := prometheus.NewPrometheusCreateParams().WithV(api.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Prometheus.PrometheusCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}

func parseLabelsFlag(labelsData []string) ([]*models.PrometheusLabelListDto, error) {
	labels := make([]*models.PrometheusLabelListDto, len(labelsData))

	for labelIndex, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, errors.New("Invalid empty billing rule label")
		}

		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid billing rule label format: %s", labelData)
		}

		labels[labelIndex] = &models.PrometheusLabelListDto{
			Label: tokens[0],
			Value: tokens[1],
		}
	}

	return labels, nil
}
