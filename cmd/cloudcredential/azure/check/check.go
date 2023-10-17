package check

import (
	"context"
	"fmt"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
	"strings"
)

type CheckOptions struct {
	AzureClientId     string
	AzureClientSecret string
	AzureTenantId     string
}

func NewCmdCheck() *cobra.Command {
	var opts CheckOptions

	cmd := &cobra.Command{
		Use:   "check",
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
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CheckAzureCommand{
		AzureClientId:     *taikuncore.NewNullableString(&opts.AzureClientId),
		AzureClientSecret: *taikuncore.NewNullableString(&opts.AzureClientSecret),
		AzureTenantId:     *taikuncore.NewNullableString(&opts.AzureTenantId),
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.CheckerAPI.CheckerAzure(context.TODO()).CheckAzureCommand(body)
	response, err := myRequest.Execute()

	if err == nil {
		out.PrintCheckSuccess("Azure cloud credential")
	}

	// Did it fail because the request failed (eg cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		myError := tk.CreateError(response, err)
		myStringError := fmt.Sprint(myError)
		if strings.Contains(myStringError, "Failed to validate") {
			err = cmderr.ErrCheckFailure("Azure cloud credential") // Taikun responded that credentials are not valid.
		} else {
			err = tk.CreateError(response, err) // Something else happened
		}

		return
	}

	return
	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		body := models.CheckAzureCommand{
			AzureClientID:     opts.AzureClientId,
			AzureClientSecret: opts.AzureClientSecret,
			AzureTenantID:     opts.AzureTenantId,
		}

		params := checker.NewCheckerAzureParams().WithV(taikungoclient.Version).WithBody(&body)

		_, err = apiClient.Client.Checker.CheckerAzure(params, apiClient)
		if err == nil {
			out.PrintCheckSuccess("Azure cloud credential")
		} else if _, isValidationProblem := err.(*checker.CheckerAzureBadRequest); isValidationProblem {
			return cmderr.ErrCheckFailure("Azure cloud credential")
		}

		return
	*/
}
