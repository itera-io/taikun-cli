package list

import (
	"context"
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"
	tk "github.com/itera-io/taikungoclient"
	taikuncore "github.com/itera-io/taikungoclient/client"
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
			"STORAGE", "storage",
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
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Proxmox cloud credentials",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(cmd, "cloud-credentials", listFields)
	cmdutils.AddColumnsFlag(cmd, listFields)

	return cmd
}

func listRun(opts *ListOptions) error {
	amazonCloudCredentials, err := ListCloudCredentialsProxmox(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(amazonCloudCredentials, listFields)
}

func ListCloudCredentialsProxmox(opts *ListOptions) (credentials []interface{}, err error) {
	myApiClient := tk.NewClient()
	//myRequest := myApiClient.Client.CloudCredentialAPI.CloudcredentialsDashboardList(context.TODO())
	myRequest := myApiClient.Client.ProxmoxCloudCredentialAPI.ProxmoxList(context.TODO())

	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	//var proxmoxCloudCredentials = make([]taikuncore.AmazonCredentialsListDto, 0)
	var proxmoxCloudCredentials = make([]taikuncore.ProxmoxListDto, 0)

	for {
		data, response, newError := myRequest.Execute()
		if newError != nil {
			err = tk.CreateError(response, err)
			return
		}

		proxmoxCloudCredentials = append(proxmoxCloudCredentials, data.Data...)

		count := int32(len(proxmoxCloudCredentials))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(proxmoxCloudCredentials)) > opts.Limit {
		proxmoxCloudCredentials = proxmoxCloudCredentials[:opts.Limit]
	}

	credentials = make([]interface{}, len(proxmoxCloudCredentials))
	for i, credential := range proxmoxCloudCredentials {
		credentials[i] = credential
	}

	return credentials, nil

}
