package add

import (
	"context"
	tk "github.com/Smidra/taikungoclient"
	taikuncore "github.com/Smidra/taikungoclient/client"
	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/complete"
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
			"NAME", "cloudCredentialName",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewVisible(
			"REGION", "awsRegion",
		),
		field.NewVisible(
			"AZ-COUNT", "awsAzCount",
		),
		field.NewHidden(
			"ACCESS-KEY-ID", "awsAccessKeyId",
		),
		field.NewHidden(
			"SECRET-ACCESS-KEY", "awsSecretAccessKey",
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
	},
)

type AddOptions struct {
	Name               string
	AWSSecretAccessKey string
	AWSAccessKeyID     string
	AWSRegion          string
	AWSAzCount         int32
	OrganizationID     int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := cobra.Command{
		Use:   "add <name>",
		Short: "Add an AWS cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AWSSecretAccessKey, "secret-access-key", "s", "", "AWS Secret Access Key (required)")
	cmdutils.MarkFlagRequired(&cmd, "secret-access-key")

	cmd.Flags().StringVarP(&opts.AWSAccessKeyID, "access-key-id", "a", "", "AWS Access Key ID (required)")
	cmdutils.MarkFlagRequired(&cmd, "access-key-id")

	cmd.Flags().StringVarP(&opts.AWSRegion, "region", "r", "", "AWS Region (required)")
	cmdutils.MarkFlagRequired(&cmd, "region")
	cmdutils.SetFlagCompletionFunc(&cmd, "region", complete.MakeAwsRegionCompletionFunc(&opts.AWSAccessKeyID, &opts.AWSSecretAccessKey))

	cmd.Flags().Int32VarP(&opts.AWSAzCount, "az-count", "z", 0, "AWS Az Count (required)")
	cmdutils.MarkFlagRequired(&cmd, "az-count")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(&cmd)
	cmdutils.AddColumnsFlag(&cmd, addFields)

	return &cmd
}

func addRun(opts *AddOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.CreateAwsCloudCommand{
		Name:               *taikuncore.NewNullableString(&opts.Name),
		AwsSecretAccessKey: *taikuncore.NewNullableString(&opts.AWSSecretAccessKey),
		AwsAccessKeyId:     *taikuncore.NewNullableString(&opts.AWSAccessKeyID),
		AzCount:            &opts.AWSAzCount,
		AwsRegion:          *taikuncore.NewNullableString(&opts.AWSRegion),
		OrganizationId:     *taikuncore.NewNullableInt32(&opts.OrganizationID),
	}

	// Execute a query into the API + graceful exit
	data, response, err := myApiClient.Client.AWSCloudCredentialAPI.AwsCreate(context.TODO()).CreateAwsCloudCommand(body).Execute()
	if err != nil {
		err = tk.CreateError(response, err)
		return
	}
	return out.PrintResult(data, addFields)

	/*
			apiClient, err := taikungoclient.NewClient()
			if err != nil {
				return
			}

			body := &models.CreateAwsCloudCommand{
				Name:                opts.Name,
				AwsSecretAccessKey:  opts.AWSSecretAccessKey,
				AwsAccessKeyID:      opts.AWSAccessKeyID,
				AwsRegion:           opts.AWSRegion,
				AzCount:             opts.AWSAzCount,
		                OrganizationID:      opts.OrganizationID,
			}

			params := aws.NewAwsCreateParams().WithV(taikungoclient.Version).WithBody(body)

			response, err := apiClient.Client.Aws.AwsCreate(params, apiClient)
			if err == nil {
				return out.PrintResult(response.Payload, addFields)
			}

			return
	*/
}
