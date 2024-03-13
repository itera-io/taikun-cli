package list

import (
	awslist "github.com/itera-io/taikun-cli/cmd/cloudcredential/aws/list"
	azlist "github.com/itera-io/taikun-cli/cmd/cloudcredential/azure/list"
	gcplist "github.com/itera-io/taikun-cli/cmd/cloudcredential/google/list"
	oslist "github.com/itera-io/taikun-cli/cmd/cloudcredential/openstack/list"
	proxmoxlist "github.com/itera-io/taikun-cli/cmd/cloudcredential/proxmox/list"
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

	googleOpts := gcplist.ListOptions{
		OrganizationID: opts.OrganizationID,
	}

	credentialsGoogle, err := gcplist.ListCloudCredentialsGoogle(&googleOpts)
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

	proxmoxOpts := proxmoxlist.ListOptions{
		OrganizationID: opts.OrganizationID,
	}

	credentialsProxmox, err := proxmoxlist.ListCloudCredentialsProxmox(&proxmoxOpts)
	if err != nil {
		return
	}

	return out.PrintResultsOfDifferentTypes(
		[]interface{}{
			credentialsAmazon,
			credentialsAzure,
			credentialsGoogle,
			credentialsOpenStack,
			credentialsProxmox,
		},
		[]string{
			"AWS",
			"Azure",
			"Google",
			"OpenStack",
			"Proxmox",
		},
		listFields,
	)
}
