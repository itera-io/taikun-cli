package create

import (
	"errors"
	"fmt"
	"strings"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/format"
	"github.com/itera-io/taikun-cli/utils/types"
	"github.com/itera-io/taikungoclient/client/prometheus"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	BillingCredentialID int32
	Labels              []string
	MetricName          string
	Name                string
	Price               float64
	PriceRate           int32
	Type                string
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := cobra.Command{
		Use:   "create <name>",
		Short: "Create a billing rule",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			if !types.MapContains(types.PrometheusTypes, opts.Type) {
				return types.UnknownFlagValueError(
					"type",
					opts.Type,
					types.MapKeys(types.PrometheusTypes),
				)
			}
			return createRun(&opts)
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
	cmdutils.RegisterStaticFlagCompletion(&cmd, "type", types.MapKeys(types.PrometheusTypes)...)

	cmdutils.AddOutputOnlyIDFlag(&cmd)

	return &cmd
}

func createRun(opts *CreateOptions) (err error) {
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

	params := prometheus.NewPrometheusCreateParams().WithV(apiconfig.Version)
	params = params.WithBody(&body)

	response, err := apiClient.Client.Prometheus.PrometheusCreate(params, apiClient)
	if err == nil {
		format.PrintResult(response.Payload,
			"id",
			"name",
			"metricName",
			"price",
			"type",
			"createdAt",
		)
	}

	return
}

func parseLabelsFlag(labelsData []string) ([]*models.PrometheusLabelListDto, error) {
	labels := make([]*models.PrometheusLabelListDto, len(labelsData))
	for i, labelData := range labelsData {
		if len(labelData) == 0 {
			return nil, errors.New("Invalid empty billing rule label")
		}
		tokens := strings.Split(labelData, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("Invalid billing rule label format: %s", labelData)
		}
		labels[i] = &models.PrometheusLabelListDto{
			Label: tokens[0],
			Value: tokens[1],
		}
	}
	return labels, nil
}
