package add

import (
	"os"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient/client/google_cloud"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"BILLING-ACCOUNT", "billingAccountName",
		),
		field.NewVisible(
			"BILLING-ACCOUNT-ID", "billingAccountId",
		),
		field.NewVisible(
			"FOLDER-ID", "folderId",
		),
		field.NewVisible(
			"REGION", "region",
		),
		field.NewVisible(
			"ZONE", "zone",
		),
	},
)

type AddOptions struct {
	BillingAccountID string
	ConfigFilePath   string
	FolderID         string
	Name             string
	OrganizationID   int32
	Region           string
	Zone             string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a Google Cloud Platform credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.BillingAccountID, "billing-account-id", "b", "", "Billing account ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "billing-account-id")

	cmd.Flags().StringVarP(&opts.ConfigFilePath, "config-file", "c", "", "Config file path (required)")
	cmdutils.MarkFlagRequired(&cmd, "config-file")

	cmd.Flags().StringVarP(&opts.FolderID, "folder-id", "f", "", "Folder ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "folder-id")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmd.Flags().StringVarP(&opts.Region, "region", "r", "", "Region (required)")
	cmdutils.MarkFlagRequired(&cmd, "region")

	cmd.Flags().StringVarP(&opts.Zone, "zone", "z", "", "Zone (required)")
	cmdutils.MarkFlagRequired(&cmd, "zone")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	configFile, err := os.Open(opts.ConfigFilePath)
	if err != nil {
		return err
	}

	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	if opts.OrganizationID == 0 {
		opts.OrganizationID, err = organization.GetDefaultOrganizationID()
		if err != nil {
			return
		}
	}

	params := google_cloud.NewGoogleCloudCreateParams().WithV(api.Version)
	params = params.WithBillingAccountID(&opts.BillingAccountID)
	params = params.WithConfig(configFile)
	params = params.WithFolderID(&opts.FolderID)
	params = params.WithName(&opts.Name)
	params = params.WithOrganizationID(&opts.OrganizationID)
	params = params.WithRegion(&opts.Region).WithZone(&opts.Zone)

	response, err := apiClient.Client.GoogleCloud.GoogleCloudCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
