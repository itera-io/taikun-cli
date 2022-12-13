package add

import (
	"errors"
	"os"

	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/cmd/organization"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient"
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
                        "AZ-COUNT", "azCount",
                ),
	},
)

type AddOptions struct {
	BillingAccountID string
	ConfigFilePath   string
	FolderID         string
	ImportProject    bool
	Name             string
	OrganizationID   int32
	Region           string
        AzCount          int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a Google Cloud Platform credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			opts.Name = args[0]
			if opts.BillingAccountID != "" {
				if opts.ImportProject {
					return cmderr.MutuallyExclusiveFlagsError("--import-project", "--billing-account-id")
				}
			} else if !opts.ImportProject {
				return errors.New("Must set --billing-acount-id if not importing a project")
			}
			if opts.FolderID != "" {
				if opts.ImportProject {
					return cmderr.MutuallyExclusiveFlagsError("--import-project", "--folder-id")
				}
			} else if !opts.ImportProject {
				return errors.New("Must set --folder-id if not importing a project")
			}
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.BillingAccountID, "billing-account-id", "b", "", "Billing account ID")

	cmd.Flags().StringVarP(&opts.ConfigFilePath, "config-file", "c", "", "Config file path (required)")
	cmdutils.MarkFlagRequired(&cmd, "config-file")

	cmd.Flags().StringVarP(&opts.FolderID, "folder-id", "f", "", "Folder ID")

	cmd.Flags().BoolVarP(&opts.ImportProject, "import-project", "i", false, "Import project")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmd.Flags().StringVarP(&opts.Region, "region", "r", "", "Region (required)")
	cmdutils.MarkFlagRequired(&cmd, "region")

        cmd.Flags().Int32VarP(&opts.AzCount, "az-count", "z", 0, "Az Count (required)")
        cmdutils.MarkFlagRequired(&cmd, "az-count")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	configFile, err := os.Open(opts.ConfigFilePath)
	if err != nil {
		return err
	}

	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	if opts.OrganizationID == 0 {
		opts.OrganizationID, err = organization.GetDefaultOrganizationID()
		if err != nil {
			return
		}
	}

	params := google_cloud.NewGoogleCloudCreateParams().WithV(taikungoclient.Version)
	params = params.WithConfig(configFile)
	params = params.WithName(&opts.Name)
	params = params.WithOrganizationID(&opts.OrganizationID)
        params = params.WithRegion(&opts.Region).WithAzCount(&opts.AzCount)

	params = params.WithImportProject(&opts.ImportProject)
	if !opts.ImportProject {
		params = params.WithBillingAccountID(&opts.BillingAccountID)
		params = params.WithFolderID(&opts.FolderID)
	}

	response, err := apiClient.Client.GoogleCloud.GoogleCloudCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
