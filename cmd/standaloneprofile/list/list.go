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
			"PUBLIC-KEY", "publicKey",
		),
		field.NewVisibleWithToStringFunc(
			"LOCK", "isLocked", out.FormatLockStatus,
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
		Short: "List standalone profiles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(&cmd, &opts.Limit)
	cmdutils.AddColumnsFlag(&cmd, listFields)
	cmdutils.AddSortByAndReverseFlags(&cmd, "standalone-profiles", listFields)

	return &cmd
}

func listRun(opts *ListOptions) (err error) {
	myApiClient := tk.NewClient()
	myRequest := myApiClient.Client.StandaloneProfileAPI.StandaloneprofileList(context.TODO())
	if opts.OrganizationID != 0 {
		myRequest = myRequest.OrganizationId(opts.OrganizationID)
	}
	if config.SortBy != "" {
		myRequest = myRequest.SortBy(config.SortBy).SortDirection(*api.GetSortDirection())
	}

	standAloneProfiles := make([]taikuncore.StandAloneProfilesListDto, 0)
	for {
		data, response, err := myRequest.Execute()
		if err != nil {
			return tk.CreateError(response, err)
		}

		standAloneProfiles = append(standAloneProfiles, data.GetData()...)

		count := int32(len(standAloneProfiles))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}

		if count == data.GetTotalCount() {
			break
		}

		myRequest = myRequest.Offset(count)
	}

	if opts.Limit != 0 && int32(len(standAloneProfiles)) > opts.Limit {
		standAloneProfiles = standAloneProfiles[:opts.Limit]
	}

	return out.PrintResults(standAloneProfiles, listFields)

	/*
		apiClient, err := taikungoclient.NewClient()
		if err != nil {
			return
		}

		params := stand_alone_profile.NewStandAloneProfileListParams().WithV(taikungoclient.Version)
		if opts.OrganizationID != 0 {
			params = params.WithOrganizationID(&opts.OrganizationID)
		}

		if config.SortBy != "" {
			params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
		}

		standAloneProfiles := make([]*models.StandAloneProfilesListDto, 0)

		for {
			response, err := apiClient.Client.StandAloneProfile.StandAloneProfileList(params, apiClient)
			if err != nil {
				return err
			}

			standAloneProfiles = append(standAloneProfiles, response.Payload.Data...)

			count := int32(len(standAloneProfiles))
			if opts.Limit != 0 && count >= opts.Limit {
				break
			}

			if count == response.Payload.TotalCount {
				break
			}

			params = params.WithOffset(&count)
		}

		if opts.Limit != 0 && int32(len(standAloneProfiles)) > opts.Limit {
			standAloneProfiles = standAloneProfiles[:opts.Limit]
		}

		return out.PrintResults(standAloneProfiles, listFields)
	*/
}
