package list

import (
	awslist "github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/list"
	azlist "github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/list"
	oslist "github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/list"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	"github.com/spf13/cobra"
)

var listFields = fields.New(
	[]*field.Field{
		field.NewVisible(
			"ID", "id",
		),
		field.NewVisible(
			"NAME", "name",
		),
		field.NewVisible(
			"ORG", "organizationName",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
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

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List all cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByAndReverseFlags(&cmd, "cloud-credentials", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
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

	return out.PrintResultsOfDifferentTypes(
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
		listFields,
	)
}
