package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/showback"
	"github.com/itera-io/taikungoclient/models"
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
		field.NewHidden(
			"PASSWORD", "password",
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

	cmdutils.AddLimitFlag(&cmd)
	cmdutils.AddSortByAndReverseFlags(&cmd, "showback-credentials", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := showback.NewShowbackCredentialsListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var showbackCredentials = make([]*models.ShowbackCredentialsListDto, 0)
	for {
		response, err := apiClient.Client.Showback.ShowbackCredentialsList(params, apiClient)
		if err != nil {
			return err
		}
		showbackCredentials = append(showbackCredentials, response.Payload.Data...)
		showbackCredentialsCount := int32(len(showbackCredentials))
		if config.Limit != 0 && showbackCredentialsCount >= config.Limit {
			break
		}
		if showbackCredentialsCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&showbackCredentialsCount)
	}

	if config.Limit != 0 && int32(len(showbackCredentials)) > config.Limit {
		showbackCredentials = showbackCredentials[:config.Limit]
	}

	out.PrintResults(showbackCredentials, listFields)
	return
}
