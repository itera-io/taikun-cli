package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/ops_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	Name               string
	PrometheusUsername string
	PrometheusPassword string
	PrometheusURL      string
	OrganizationID     int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add a billing credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
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

func addRun(opts *AddOptions) (err error) {
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

	params := ops_credentials.NewOpsCredentialsCreateParams().WithV(api.Version).WithBody(body)
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