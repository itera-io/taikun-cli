package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/ops_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Name               string
	PrometheusUsername string
	PrometheusPassword string
	PrometheusURL      string
	OrganizationID     int32
	IDOnly             bool
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a billing credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.PrometheusUsername, "prometheus-username", "u", "", "Prometheus Username (required)")
	cmdutils.MarkFlagRequired(cmd, "prometheus-username")

	cmd.Flags().StringVarP(&opts.PrometheusPassword, "prometheus-password", "p", "", "Prometheus Password (required)")
	cmdutils.MarkFlagRequired(cmd, "prometheus-password")

	cmd.Flags().StringVar(&opts.PrometheusURL, "prometheus-url", "", "Prometheus URL (required)")
	cmdutils.MarkFlagRequired(cmd, "prometheus-url")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddIdOnlyFlag(cmd, &opts.IDOnly)

	return cmd
}

func printResult(resource interface{}) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(resource)
	} else if config.OutputFormat == config.OutputFormatTable {
		format.PrettyPrintApiResponseTable(resource,
			"id",
			"name",
			"organizationName",
			"prometheusUsername",
			"prometheusUrl",
			"isDefault",
			"isLocked",
		)
	}
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.OperationCredentialsCreateCommand{
		Name:               opts.Name,
		PrometheusUsername: opts.PrometheusUsername,
		PrometheusPassword: opts.PrometheusPassword,
		PrometheusURL:      opts.PrometheusURL,
		OrganizationID:     opts.OrganizationID,
	}

	params := ops_credentials.NewOpsCredentialsCreateParams().WithV(apiconfig.Version).WithBody(body)
	response, err := apiClient.Client.OpsCredentials.OpsCredentialsCreate(params, apiClient)
	if err == nil {
		if opts.IDOnly {
			format.PrintResourceID(response.Payload)
		} else {
			printResult(response.Payload)
		}
	}

	return
}
