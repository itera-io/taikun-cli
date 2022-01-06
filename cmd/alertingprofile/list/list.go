package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID       int32
	ReverseSortDirection bool
	SortBy               string
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List alerting profiles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByFlag(cmd, &opts.SortBy, models.AlertingProfilesListDto{})

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_profiles.NewAlertingProfilesListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var alertingProfiles = make([]*models.AlertingProfilesListDto, 0)
	for {
		response, err := apiClient.Client.AlertingProfiles.AlertingProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		alertingProfiles = append(alertingProfiles, response.Payload.Data...)
		alertingProfilesCount := int32(len(alertingProfiles))
		if config.Limit != 0 && alertingProfilesCount >= config.Limit {
			break
		}
		if alertingProfilesCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&alertingProfilesCount)
	}

	if config.Limit != 0 && int32(len(alertingProfiles)) > config.Limit {
		alertingProfiles = alertingProfiles[:config.Limit]
	}

	format.PrintResults(alertingProfiles,
		"id",
		"name",
		"organizationName",
		"slackConfigurationName",
		"reminder",
		"isLocked",
	)
	return
}
