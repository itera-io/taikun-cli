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
	taikunshowback "github.com/itera-io/taikungoclient/showbackclient"
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
			"URL", "url",
		),
		field.NewVisible(
			"USERNAME", "username",
		),
		field.NewVisible(
			"LOCK", "isLocked",
		),
		field.NewVisibleWithToStringFunc(
			"CREATED-AT", "createdAt", out.FormatDateTimeString,
		),
		field.NewVisible(
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

	cmd := cobra.Command{
		Use:   "list",
		Short: "List showback credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "showback-credentials", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	// Create and authenticated client to the Taikun API
	myApiClient := tk.NewClient()

	// Prepare the arguments for the query
	myRequest := myApiClient.ShowbackClient.ShowbackCredentialsAPI.ShowbackcredentialsList(context.TODO())
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}

	var showbackCredentials = make([]taikunshowback.ShowbackCredentialsListDto, 0)
	// Send the query and move offset until you get all the data
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		showbackCredentials = append(showbackCredentials, data.GetData()...)
		showbackCredentialsCount := int32(len(showbackCredentials))

		if opts.Limit != 0 && showbackCredentialsCount >= opts.Limit {
			break
		}

		if showbackCredentialsCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(showbackCredentialsCount)
	}

	if opts.Limit != 0 && int32(len(showbackCredentials)) > opts.Limit {
		showbackCredentials = showbackCredentials[:opts.Limit]
	}

	return out.PrintResults(showbackCredentials, listFields)

}
