package list

import (
	"taikun-cli/api"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/ops_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit          int32
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List billing credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return utils.NegativeLimitFlagError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	return cmd
}

func printResults(billingCredentials []*models.OperationCredentialsListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		utils.PrettyPrintJson(billingCredentials)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(billingCredentials))
		for i, billingCredential := range billingCredentials {
			data[i] = billingCredential
		}
		utils.PrettyPrintTable(data,
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

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := ops_credentials.NewOpsCredentialsListParams().WithV(utils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}

	var billingCredentials = make([]*models.OperationCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.OpsCredentials.OpsCredentialsList(params, apiClient)
		if err != nil {
			return err
		}
		billingCredentials = append(billingCredentials, response.Payload.Data...)
		count := int32(len(billingCredentials))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(billingCredentials)) > opts.Limit {
		billingCredentials = billingCredentials[:opts.Limit]
	}

	printResults(billingCredentials)
	return
}
