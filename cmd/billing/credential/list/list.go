package list

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

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List billing credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmdutils.AddLimitFlag(cmd)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := ops_credentials.NewOpsCredentialsListParams().WithV(apiconfig.Version)
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
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(billingCredentials)) > config.Limit {
		billingCredentials = billingCredentials[:config.Limit]
	}

	format.PrintResults(billingCredentials,
		"id",
		"name",
		"organizationName",
		"prometheusUsername",
		"prometheusUrl",
		"isDefault",
		"isLocked",
	)
	return
}
