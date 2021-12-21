package create

import (
	"fmt"

	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/checker"
	"github.com/itera-io/taikungoclient/client/s3_credentials"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	OrganizationID int32
	S3Name         string
	S3AccessKey    string
	S3SecretKey    string
	S3Endpoint     string
	S3Region       string
	IDOnly         bool
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create a backup credential",
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
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.S3AccessKey, "s3-access-key", "a", "", "S3 access key (required)")
	cmdutils.MarkFlagRequired(cmd, "s3-access-key")

	cmd.Flags().StringVarP(&opts.S3SecretKey, "s3-secret-key", "s", "", "S3 secret key (required)")
	cmdutils.MarkFlagRequired(cmd, "s3-secret-key")

	cmd.Flags().StringVarP(&opts.S3Endpoint, "s3-endpoint", "e", "", "S3 endpoint (required)")
	cmdutils.MarkFlagRequired(cmd, "s3-endpoint")

	cmd.Flags().StringVarP(&opts.S3Region, "s3-region", "r", "", "S3 region (required)")
	cmdutils.MarkFlagRequired(cmd, "s3-region")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddIdOnlyFlag(cmd, &opts.IDOnly)

	return cmd
}

func backupCredentialIsValid(opts *CreateOptions) (bool, error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return false, err
	}
	body := models.CheckS3Command{
		S3AccessKeyID: opts.S3AccessKey,
		S3SecretKey:   opts.S3SecretKey,
		S3Endpoint:    opts.S3Endpoint,
		S3Region:      opts.S3Region,
	}
	params := checker.NewCheckerS3Params().WithV(apiconfig.Version).WithBody(&body)
	_, err = apiClient.Client.Checker.CheckerS3(params, apiClient)
	return err == nil, nil
}

func printResult(resource interface{}) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(resource)
	} else if config.OutputFormat == config.OutputFormatTable {
		format.PrettyPrintApiResponseTable(resource,
			"id",
			"organizationName",
			"s3Name",
			"s3AccessKeyId",
			"s3Endpoint",
			"s3Region",
			"isDefault",
			"isLocked",
		)
	}
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
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

	params := s3_credentials.NewS3CredentialsCreateParams().WithV(apiconfig.Version).WithBody(&body)
	response, err := apiClient.Client.S3Credentials.S3CredentialsCreate(params, apiClient)
	if err == nil {
		if opts.IDOnly {
			format.PrintResourceID(response.Payload)
		} else {
			printResult(response.Payload)
		}
	}

	return
}
