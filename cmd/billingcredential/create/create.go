package create

import (
	"taikun-cli/api"
	"taikun-cli/utils"

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
	utils.MarkFlagRequired(cmd, "prometheus-username")

	cmd.Flags().StringVarP(&opts.PrometheusPassword, "prometheus-password", "p", "", "Prometheus Password (required)")
	utils.MarkFlagRequired(cmd, "prometheus-password")

	cmd.Flags().StringVar(&opts.PrometheusURL, "prometheus-url", "", "Prometheus URL (required)")
	utils.MarkFlagRequired(cmd, "prometheus-url")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	return cmd
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

	params := ops_credentials.NewOpsCredentialsCreateParams().WithV(utils.ApiVersion).WithBody(body)
	response, err := apiClient.Client.OpsCredentials.OpsCredentialsCreate(params, apiClient)
	if err == nil {
		utils.PrettyPrintJson(response.Payload)
	}

	return
}
