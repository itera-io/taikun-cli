package check

import (
	"taikun-cli/api"
	"taikun-cli/apiconfig"
	"taikun-cli/cmd/cmderr"
	"taikun-cli/cmd/cmdutils"
	"taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CheckOptions struct {
	AzureClientId     string
	AzureClientSecret string
	AzureTenantId     string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := &cobra.Command{
		Use:   "check <name>",
		Short: "Check the validity of an Azure cloud credential",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return checkRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AzureClientId, "client-id", "c", "", "Azure Client ID (required)")
	cmdutils.MarkFlagRequired(cmd, "client-id")

	cmd.Flags().StringVarP(&opts.AzureClientSecret, "client-secret", "p", "", "Azure Client Secret (required)")
	cmdutils.MarkFlagRequired(cmd, "client-secret")

	cmd.Flags().StringVarP(&opts.AzureTenantId, "tenant-id", "t", "", "Azure Tenant ID (required)")
	cmdutils.MarkFlagRequired(cmd, "tenant-id")

	return cmd
}

func checkRun(opts *CheckOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := models.CheckAzureCommand{
		AzureClientID:     opts.AzureClientId,
		AzureClientSecret: opts.AzureClientSecret,
		AzureTenantID:     opts.AzureTenantId,
	}

	params := checker.NewCheckerAzureParams().WithV(apiconfig.Version).WithBody(&body)
	_, err = apiClient.Client.Checker.CheckerAzure(params, apiClient)
	if err == nil {
		format.PrintCheckSuccess("Azure cloud credential")
	} else if _, isValidationProblem := err.(*checker.CheckerAzureBadRequest); isValidationProblem {
		return cmderr.CheckFailureError("Azure cloud credential")
	}

	return
}