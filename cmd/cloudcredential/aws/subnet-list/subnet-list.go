package subnet_list

import (
	"context"

	"github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/complete"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewHidden(
			"SUBNET-ID", "subnetId",
		),
		field.NewVisible(
			"ZONE", "zone",
		),
		field.NewVisible(
			"STATE", "state",
		),
		field.NewHidden(
			"HAS-IP-V6", "hasIpv6",
		),
		field.NewVisible(
			"CIDR", "cidr",
		),
	},
)

type SubnetListOptions struct {
	AWSSecretAccessKey string
	AWSAccessKeyID     string
	AWSRegion          string
	VPCID              string
}

func NewCmdSubnetList() *cobra.Command {
	var opts SubnetListOptions

	cmd := &cobra.Command{
		Use:   "subnet-list",
		Short: "List subnets for an AWS cloud credential",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return subnetList(&opts)
		},
	}

	cmd.Flags().StringVarP(&opts.AWSSecretAccessKey, "secret-access-key", "s", "", "AWS Secret Access Key (required)")
	cmdutils.MarkFlagRequired(cmd, "secret-access-key")

	cmd.Flags().StringVarP(&opts.AWSAccessKeyID, "access-key-id", "a", "", "AWS Access Key ID (required)")
	cmdutils.MarkFlagRequired(cmd, "access-key-id")

	cmd.Flags().StringVarP(&opts.AWSRegion, "region", "r", "", "AWS Region (required)")
	cmdutils.MarkFlagRequired(cmd, "region")
	cmdutils.SetFlagCompletionFunc(cmd, "region", complete.MakeAwsRegionCompletionFunc(&opts.AWSAccessKeyID, &opts.AWSSecretAccessKey))

	cmd.Flags().StringVarP(&opts.VPCID, "vpc-id", "", "", "VPC ID (required)")
	cmdutils.MarkFlagRequired(cmd, "vpc-id")

	return cmd
}

func subnetList(opts *SubnetListOptions) error {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	body := taikuncore.AwsSubnetListCommand{
		AwsAccessKeyId:     *taikuncore.NewNullableString(&opts.AWSAccessKeyID),
		AwsSecretAccessKey: *taikuncore.NewNullableString(&opts.AWSSecretAccessKey),
		AwsRegion:          *taikuncore.NewNullableString(&opts.AWSRegion),
		VpcId:              *taikuncore.NewNullableString(&opts.VPCID),
	}

	// Execute a query into the API + graceful exit
	myRequest := myApiClient.Client.AWSCloudCredentialAPI.AwsSubnetList(context.TODO()).AwsSubnetListCommand(body)
	subnets, response, err := myRequest.Execute()
	// Did it fail because the request failed (e.g. cannot connect to Taikun) or because the credentials are not valid?
	if err != nil {
		return tk.CreateError(response, err)
	}

	return out.PrintResults(subnets, listFields)
}
