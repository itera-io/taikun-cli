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
			"BILLING-ACCOUNT", "billingAccountName",
		),
		field.NewVisible(
			"FOLDER-ID", "folderId",
		),
		field.NewVisible(
			"PROJECT-ID", "projectId",
		),
		field.NewVisible(
			"PARTNER", "partnerName",
		),
		field.NewVisible(
			"REGION", "region",
		),
		field.NewVisible(
			"ZONES", "zones",
		),
		field.NewHiddenWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHidden(
			"DEFAULT", "isDefault",
		),
	},
)

type ListOptions struct {
	OrganizationID int32
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List Google Cloud Platform credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "cloud-credentials", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	googleCloudCredentials, err := ListCloudCredentialsGoogle(opts)
	if err != nil {
		return err
	}

	return out.PrintResults(googleCloudCredentials, listFields)
}

func ListCloudCredentialsGoogle(opts *ListOptions) (credentials []interface{}, err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.GoogleAPI.GooglecloudList(context.TODO())
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var googlecloudCloudCredentials = make([]taikuncore.GoogleCredentialsListDto, 0)

	for {
		data, response, newError := myRequest.Execute()
		if newError != nil {
			err = tk.CreateError(response, err)
			return
		}

		googlecloudCloudCredentials = append(googlecloudCloudCredentials, data.GetData()...)

		count := int32(len(googlecloudCloudCredentials))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(googlecloudCloudCredentials)) > opts.Limit {
		googlecloudCloudCredentials = googlecloudCloudCredentials[:opts.Limit]
	}

	credentials = make([]interface{}, len(googlecloudCloudCredentials))
	for i, credential := range googlecloudCloudCredentials {
		credentials[i] = credential
	}

	return credentials, nil

}
