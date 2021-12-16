package list

import (
	"taikun-cli/api"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/alerting_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	Limit                int32
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
			if opts.Limit < 0 {
				return utils.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			return listRun(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

	return cmd
}

func printResults(alertingProfiles []*models.AlertingProfilesListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		utils.PrettyPrintJson(alertingProfiles)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(alertingProfiles))
		for i, alertingProfile := range alertingProfiles {
			data[i] = alertingProfile
		}
		utils.PrettyPrintTable(data,
			"id",
			"name",
			"organizationName",
			"slackConfigurationName",
			"reminder",
			"isLocked",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := alerting_profiles.NewAlertingProfilesListParams().WithV(utils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		utils.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&utils.SortDirection)
	}

	var alertingProfiles = make([]*models.AlertingProfilesListDto, 0)
	for {
		response, err := apiClient.Client.AlertingProfiles.AlertingProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		alertingProfiles = append(alertingProfiles, response.Payload.Data...)
		alertingProfilesCount := int32(len(alertingProfiles))
		if opts.Limit != 0 && alertingProfilesCount >= opts.Limit {
			break
		}
		if alertingProfilesCount == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&alertingProfilesCount)
	}

	if opts.Limit != 0 && int32(len(alertingProfiles)) > opts.Limit {
		alertingProfiles = alertingProfiles[:opts.Limit]
	}

	printResults(alertingProfiles)
	return
}
