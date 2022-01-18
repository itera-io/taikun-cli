package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/utils/out"
	"github.com/itera-io/taikun-cli/utils/out/field"
	"github.com/itera-io/taikun-cli/utils/out/fieldnames"
	"github.com/itera-io/taikun-cli/utils/out/fields"

	"github.com/itera-io/taikungoclient/client/access_profiles"
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
		field.NewVisible(
			"HTTP-PROXY", "httpProxy",
		),
		field.NewVisibleWithToStringFunc(
			fieldnames.IsLocked, "isLocked", out.FormatLockStatus,
		),
		field.NewVisibleWithToStringFunc(
			"LAST-MODIFIED", "lastModified", out.FormatDateTimeString,
		),
		field.NewHidden(
			"CREATED-BY", "createdBy",
		),
		field.NewHidden(
			"LAST-MODIFIED-BY", "lastModifiedBy",
		),
		field.NewHidden(
			"ORG-ID", "organizationId",
		),
	},
)

type ListOptions struct {
	OrganizationID int32

	Limit int32

	SortBy               string
	ReverseSortDirection bool

	Columns    []string
	AllColumns bool
}

// func (opts *ListOptions) GetLimitOption() *int32 {
// 	return &opts.Limit
// }

// func (opts *ListOptions) GetSortByOption() *string {
// 	return &opts.SortBy
// }

// func (opts *ListOptions) GetReverseSortDirectionOption() *bool {
// 	return &opts.ReverseSortDirection
// }

// func (opts *ListOptions) GetColumnsOption() []string {
// 	return opts.Columns
// }

// func (opts *ListOptions) GetAllColumnsOption() *bool {
// 	return &opts.AllColumns
// }

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List access profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts)
	cmdutils.AddSortByAndReverseFlags(&cmd, &opts, listFields)
	cmdutils.AddColumnsFlag(&cmd, &opts, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := access_profiles.NewAccessProfilesListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(opts.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var accessProfiles = make([]*models.AccessProfilesListDto, 0)
	for {
		response, err := apiClient.Client.AccessProfiles.AccessProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		accessProfiles = append(accessProfiles, response.Payload.Data...)
		count := int32(len(accessProfiles))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(accessProfiles)) > opts.Limit {
		accessProfiles = accessProfiles[:opts.Limit]
	}

	out.PrintResults(accessProfiles, listFields)
	return
}
