package add

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/aws"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type AddOptions struct {
	Name                string
	AWSSecretAccessKey  string
	AWSAccessKeyID      string
	AWSRegion           string
	AWSAvailabilityZone string
	OrganizationID      int32
}

func NewCmdAdd() *cobra.Command {
	var opts AddOptions

	cmd := &cobra.Command{
		Use:   "add <name>",
		Short: "Add an AWS cloud credential",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.Name = args[0]
			return addRun(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AWSSecretAccessKey, "secret-access-key", "s", "", "AWS Secret Access Key (required)")
	cmdutils.MarkFlagRequired(cmd, "secret-access-key")

	cmd.Flags().StringVarP(&opts.AWSAccessKeyID, "access-key-id", "a", "", "AWS Access Key ID (required)")
	cmdutils.MarkFlagRequired(cmd, "access-key-id")

	cmd.Flags().StringVarP(&opts.AWSRegion, "region", "r", "", "AWS Region (required)")
	cmdutils.MarkFlagRequired(cmd, "region")
	cmdutils.RegisterFlagCompletionFunc(cmd, "region", func(cmd *cobra.Command, args []string, toComplete string) (completions []string) {
		completions = make([]string, 0)

		apiClient, err := api.NewClient()
		if err != nil {
			return
		}

		params := aws.NewAwsRegionListParams().WithV(api.Version)
		result, err := apiClient.Client.Aws.AwsRegionList(params, apiClient)
		if err != nil {
			return
		}

		for _, region := range result.Payload {
			completions = append(completions, region.Region)
		}

		return
	})

	cmd.Flags().StringVarP(&opts.AWSAvailabilityZone, "availability-zone", "z", "", "AWS Availability Zone")
	cmdutils.MarkFlagRequired(cmd, "availability-zone")

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID")

	cmdutils.AddOutputOnlyIDFlag(cmd)

	return cmd
}

func addRun(opts *AddOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	body := &models.CreateAwsCloudCommand{
		Name:                opts.Name,
		AwsSecretAccessKey:  opts.AWSSecretAccessKey,
		AwsAccessKeyID:      opts.AWSAccessKeyID,
		AwsRegion:           opts.AWSRegion,
		AwsAvailabilityZone: opts.AWSAvailabilityZone,
		OrganizationID:      opts.OrganizationID,
	}

	params := aws.NewAwsCreateParams().WithV(api.Version).WithBody(body)
	response, err := apiClient.Client.Aws.AwsCreate(params, apiClient)
	if err == nil {
		out.PrintResult(response.Payload,
			"id",
			"cloudCredentialName",
			"organizationName",
			"awsRegion",
			"awsAvailabilityZone",
		)
	}

	return
}
