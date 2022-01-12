package list

import (
	awslist "github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/list"
	azlist "github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/list"
	oslist "github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/list"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikungoclient/models"

	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByAndReverseFlags(cmd,
		models.AmazonCredentialsListDto{},
		models.OpenstackCredentialsListDto{},
		models.AzureCredentialsListDto{},
	)

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	amazonOpts := awslist.ListOptions{
		OrganizationID: opts.OrganizationID,
	}
	credentialsAmazon, err := awslist.ListCloudCredentialsAws(&amazonOpts)
	if err != nil {
		return
	}

	azureOpts := azlist.ListOptions{
		OrganizationID: opts.OrganizationID,
	}
	credentialsAzure, err := azlist.ListCloudCredentialsAzure(&azureOpts)
	if err != nil {
		return
	}

	openstackOpts := oslist.ListOptions{
		OrganizationID: opts.OrganizationID,
	}
	credentialsOpenStack, err := oslist.ListCloudCredentialsOpenStack(&openstackOpts)
	if err != nil {
		return
	}

	out.PrintMultipleResults(
		[]interface{}{
			credentialsAmazon,
			credentialsAzure,
			credentialsOpenStack,
		},
		[]string{
			"AWS",
			"Azure",
			"OpenStack",
		},
		"id",
		"name",
		"organizationName",
		"createdBy",
		"isDefault",
		"isLocked",
	)

	return
}
