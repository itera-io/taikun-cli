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
		field.NewVisible(
			"SLACK-CONFIG", "slackConfigurationName",
		),
		field.NewVisible(
			"REMINDER", "reminder",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
		),
		field.NewHiddenWithToStringFunc(
			"LAST-MODIFIED", "lastModified", out.FormatDateTimeString,
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
	Limit          int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := cobra.Command{
		Use:   "list",
		Short: "List alerting profiles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddSortByAndReverseFlags(&cmd, "alerting-profiles", listFields)
	cmdutils.AddColumnsFlag(&cmd, listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.AlertingProfilesAPI.AlertingprofilesList(context.TODO())
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	var alertingProfiles = make([]taikuncore.AlertingProfilesListDto, 0)
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		alertingProfiles = append(alertingProfiles, data.GetData()...)

		alertingProfilesCount := int32(len(alertingProfiles))
		if opts.Limit != 0 && alertingProfilesCount >= opts.Limit {
			break
		}

		if alertingProfilesCount == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(alertingProfilesCount)
	}

	if opts.Limit != 0 && int32(len(alertingProfiles)) > opts.Limit {
		alertingProfiles = alertingProfiles[:opts.Limit]
	}

	return out.PrintResults(alertingProfiles, listFields)

}
