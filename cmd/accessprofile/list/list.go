package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/access_profiles"
	"github.com/itera-io/taikungoclient/models"
	"github.com/spf13/cobra"
)

type ListOptions struct {
	OrganizationID int32
}

func NewCmdList() *cobra.Command {
	var opts ListOptions

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List access profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByAndReverseFlags(cmd, models.AccessProfilesListDto{})

	return cmd
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
	if config.ReverseSortDirection {
		api.ReverseSortDirection()
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(&api.SortDirection)
	}

	var accessProfiles = make([]*models.AccessProfilesListDto, 0)
	for {
		response, err := apiClient.Client.AccessProfiles.AccessProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		accessProfiles = append(accessProfiles, response.Payload.Data...)
		count := int32(len(accessProfiles))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(accessProfiles)) > config.Limit {
		accessProfiles = accessProfiles[:config.Limit]
	}

	out.PrintResults(accessProfiles,
		"id",
		"name",
		"organizationName",
		"isLocked",
	)
	return
}
