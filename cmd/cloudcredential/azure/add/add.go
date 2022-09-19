package add

import (
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/azure"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "cloudCredentialName",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"LOCATION", "azureLocation",
		),
		field.NewVisible(
			"AVAILABILITY-ZONE", "azureAvailabilityZone",
		),
		field.NewHidden(
			"CLIENT-ID", "azureClientId",
		),
		field.NewHidden(
			"CLIENT-SECRET", "azureClientSecret",
		),
		field.NewHidden(
			"SUBSCRIPTION-ID", "azureSubscriptionId",
		),
		field.NewHidden(
			"TENANT-ID", "azureTenantId",
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	Name                  string
	AzureSubscriptionId   string
	AzureClientId         string
	AzureClientSecret     string
	AzureTenantId         string
	AzureLocation         string
	AzureAvailabilityZone string
	OrganizationID        int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add an Azure cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AzureSubscriptionId, "subscription-id", "s", "", "Azure Subscription ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "subscription-id")

	cmd.Flags().StringVarP(&opts.AzureClientId, "client-id", "c", "", "Azure Client ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "client-id")

	cmd.Flags().StringVarP(&opts.AzureClientSecret, "client-secret", "p", "", "Azure Client Secret (required)")
	cmdutils.MarkFlagRequired(&cmd, "client-secret")

	cmd.Flags().StringVarP(&opts.AzureTenantId, "tenant-id", "t", "", "Azure Tenant ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "tenant-id")

	cmd.Flags().StringVarP(&opts.AzureLocation, "location", "l", "", "Azure Location (required)")
	cmdutils.MarkFlagRequired(&cmd, "location")

	cmd.Flags().StringVarP(&opts.AzureAvailabilityZone, "availability-zone", "a", "", "Azure Availability Zone (required)")
	cmdutils.MarkFlagRequired(&cmd, "availability-zone")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := &models.CreateAzureCloudCommand{
		Name:                  opts.Name,
		AzureSubscriptionID:   opts.AzureSubscriptionId,
		AzureClientID:         opts.AzureClientId,
		AzureClientSecret:     opts.AzureClientSecret,
		AzureTenantID:         opts.AzureTenantId,
		AzureLocation:         opts.AzureLocation,
		AzureAvailabilityZone: opts.AzureAvailabilityZone,
		OrganizationID:        opts.OrganizationID,
	}

	params := azure.NewAzureCreateParams().WithV(taikungoclient.Version).WithBody(body)

	response, err := apiClient.Client.Azure.AzureCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
