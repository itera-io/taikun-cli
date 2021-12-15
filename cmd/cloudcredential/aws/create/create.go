package create

import (
	"taikun-cli/api"
	"taikun-cli/cmd/cmdutils"

	"github.com/itera-io/taikungoclient/client/aws"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type CreateOptions struct {
	Name                string
	AWSSecretAccessKey  string
	AWSAccessKeyID      string
	AWSRegion           string
	AWSAvailabilityZone string
	OrganizationID      int32
}

func NewCmdCreate() *cobra.Command {
	var opts CreateOptions

	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create an aws cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return createRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AWSSecretAccessKey, "secret-access-key", "s", "", "AWS Secret Access Key (required)")
	cmdutils.MarkFlagRequired(cmd, "secret-access-key")

	cmd.Flags().StringVarP(&opts.AWSAccessKeyID, "access-key-id", "a", "", "AWS Access Key ID (required)")
	cmdutils.MarkFlagRequired(cmd, "access-key-id")

	cmd.Flags().StringVarP(&opts.AWSRegion, "region", "r", "", "AWS Region (required)")
	cmdutils.MarkFlagRequired(cmd, "region")

	cmd.Flags().StringVarP(&opts.AWSAvailabilityZone, "availability-zone", "z", "", "AWS Availability Zone")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	return cmd
}

func createRun(opts *CreateOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CreateAwsCloudCommand{
		Name:                opts.Name,
		AwsSecretAccessKey:  opts.AWSSecretAccessKey,
		AwsAccessKeyID:      opts.AWSAccessKeyID,
		AwsRegion:           cmdutils.GetAWSRegion(opts.AWSRegion),
		AwsAvailabilityZone: opts.AWSAvailabilityZone,
		OrganizationID:      opts.OrganizationID,
	}

	params := aws.NewAwsCreateParams().WithV(cmdutils.ApiVersion).WithBody(body)
	response, err := apiClient.Client.Aws.AwsCreate(params, apiClient)
	if err == nil {
		cmdutils.PrettyPrint(response.Payload)
	}

	return
}
