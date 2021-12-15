package create

import (
	"taikun-cli/api"
	"taikun-cli/utils"

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
	utils.MarkFlagRequired(cmd, "secret-access-key")

	cmd.Flags().StringVarP(&opts.AWSAccessKeyID, "access-key-id", "a", "", "AWS Access Key ID (required)")
	utils.MarkFlagRequired(cmd, "access-key-id")

	cmd.Flags().StringVarP(&opts.AWSRegion, "region", "r", "", "AWS Region (required)")
	utils.MarkFlagRequired(cmd, "region")
	utils.RegisterFlagCompletionFunc(cmd, "region", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		apiClient, err := api.NewClient()
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveDefault
		}

		params := aws.NewAwsRegionListParams().WithV(utils.ApiVersion)
		result, err := apiClient.Client.Aws.AwsRegionList(params, apiClient)
		if err != nil {
			return []string{}, cobra.ShellCompDirectiveDefault
		}

		regionNames := make([]string, 0)
		for _, region := range result.Payload {
			regionNames = append(regionNames, region.Region)
		}

		return regionNames, cobra.ShellCompDirectiveDefault
	})

	cmd.Flags().StringVarP(&opts.AWSAvailabilityZone, "availability-zone", "z", "", "AWS Availability Zone")
	utils.MarkFlagRequired(cmd, "availability-zone")

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
		AwsRegion:           opts.AWSRegion,
		AwsAvailabilityZone: opts.AWSAvailabilityZone,
		OrganizationID:      opts.OrganizationID,
	}

	params := aws.NewAwsCreateParams().WithV(utils.ApiVersion).WithBody(body)
	response, err := apiClient.Client.Aws.AwsCreate(params, apiClient)
	if err == nil {
		utils.PrettyPrintJson(response.Payload)
	}

	return
}
