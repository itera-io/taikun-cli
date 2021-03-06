package add

import (
	"fmt"

	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/itera-io/taikungoclient"
	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
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
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return false, err
	}

	body := models.CheckS3Command{
		S3AccessKeyID: opts.S3AccessKey,
		S3SecretKey:   opts.S3SecretKey,
		S3Endpoint:    opts.S3Endpoint,
		S3Region:      opts.S3Region,
	}
	params := checker.NewCheckerS3Params().WithV(taikungoclient.Version).WithBody(&body)
	_, err = apiClient.Client.Checker.CheckerS3(params, apiClient)

	return err == nil, nil
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := taikungoclient.NewClient()
	if err != nil {
		return
	}

	body := models.BackupCredentialsCreateCommand{
		S3AccessKeyID: opts.S3AccessKey,
		S3Endpoint:    opts.S3Endpoint,
		S3Name:        opts.S3Name,
		S3Region:      opts.S3Region,
		S3SecretKey:   opts.S3SecretKey,
	}
	if opts.OrganizationID != 0 {
		body.OrganizationID = opts.OrganizationID
	}

	params := s3_credentials.NewS3CredentialsCreateParams().WithV(taikungoclient.Version).WithBody(&body)

	response, err := apiClient.Client.S3Credentials.S3CredentialsCreate(params, apiClient)
	if err == nil {
		return out.PrintResult(response.Payload, addFields)
	}

	return
}
