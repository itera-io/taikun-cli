package create

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

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

	cmd.Flags().StringVarP(&opts.PrometheusUsername, "username", "l", "", "Prometheus Username (required)")
	cmdutils.MarkFlagRequired(cmd, "username")

	cmd.Flags().StringVarP(&opts.PrometheusPassword, "password", "p", "", "Prometheus Password (required)")
	cmdutils.MarkFlagRequired(cmd, "password")

	cmd.Flags().StringVarP(&opts.PrometheusURL, "url", "u", "", "Prometheus URL (required)")
	cmdutils.MarkFlagRequired(cmd, "url")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(cmd)

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

	params := ops_credentials.NewOpsCredentialsCreateParams().WithV(apiconfig.Version).WithBody(body)
	response, err := apiClient.Client.OpsCredentials.OpsCredentialsCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"name",
			"organizationName",
			"prometheusUsername",
			"prometheusUrl",
			"isDefault",
			"isLocked",
		)
	}

	return
}
