package add

import (
	"context"
	"fmt"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

var addFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
		field.NewVisible(
			"S3-NAME", "s3Name",
		),
		field.NewVisible(
			"S3-ACCESS-KEY-ID", "s3AccessKeyId",
		),
		field.NewVisible(
			"S3-ENDPOINT", "s3Endpoint",
		),
		field.NewVisible(
			"S3-REGION", "s3Region",
		),
		field.NewVisible(
			"DEFAULT", "isDefault",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	OrganizationID int32
	S3Name         string
	S3AccessKey    string
	S3SecretKey    string
	S3Endpoint     string
	S3Region       string
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add a backup credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			isValid, err := backupCredentialIsValid(&opts)
			if err != nil {
				return err
			}
			if !isValid {
				return fmt.Errorf("backup credential must be valid")
			}
			opts.S3Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.S3AccessKey, "s3-access-key", "a", "", "S3 access key (required)")
	cmdutils.MarkFlagRequired(&cmd, "s3-access-key")

	cmd.Flags().StringVarP(&opts.S3SecretKey, "s3-secret-key", "s", "", "S3 secret key (required)")
	cmdutils.MarkFlagRequired(&cmd, "s3-secret-key")

	cmd.Flags().StringVarP(&opts.S3Endpoint, "s3-endpoint", "e", "", "S3 endpoint (required)")
	cmdutils.MarkFlagRequired(&cmd, "s3-endpoint")

	cmd.Flags().StringVarP(&opts.S3Region, "s3-region", "r", "", "S3 region (required)")
	cmdutils.MarkFlagRequired(&cmd, "s3-region")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func backupCredentialIsValid(opts *AddOptions) (bool, error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CheckS3Command{
		S3AccessKeyId: *taikuncore.NewNullableString(&opts.S3AccessKey),
		S3SecretKey:   *taikuncore.NewNullableString(&opts.S3SecretKey),
		S3Endpoint:    *taikuncore.NewNullableString(&opts.S3Endpoint),
		S3Region:      *taikuncore.NewNullableString(&opts.S3Region),
	}

	// Execute a query into the API + graceful exit
	response, err := myApiClient.Client.CheckerAPI.CheckerS3(context.TODO()).CheckS3Command(body).Execute()
	if err != nil {
		return false, tk.CreateError(response, err)
	}

	return true, nil

}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.BackupCredentialsCreateCommand{
		S3Name:         *taikuncore.NewNullableString(&opts.S3Name),
		S3AccessKeyId:  *taikuncore.NewNullableString(&opts.S3AccessKey),
		S3SecretKey:    *taikuncore.NewNullableString(&opts.S3SecretKey),
		S3Endpoint:     *taikuncore.NewNullableString(&opts.S3Endpoint),
		S3Region:       *taikuncore.NewNullableString(&opts.S3Region),
		OrganizationId: *taikuncore.NewNullableInt32(&opts.OrganizationID),
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.S3CredentialsAPI.S3credentialsCreate(context.TODO()).BackupCredentialsCreateCommand(body).Execute()
	if err != nil {
		return tk.CreateError(response, err)
	}

	// Manipulate the gathered data
	return out.PrintResult(data, addFields)

}
