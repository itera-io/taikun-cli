package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/apiconfig"
	"github.com/itera-io/taikun-cli/cmd/cmderr"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/format"

	"github.com/itera-io/taikungoclient/client/opa_profiles"
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
		Short: "List policy profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return cmderr.NegativeLimitFlagError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddSortByFlag(cmd, &opts.SortBy, models.OpaProfileListDto{})

	return cmd
}

func printResults(policyProfiles []*models.OpaProfileListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		format.PrettyPrintJson(policyProfiles)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(policyProfiles))
		for i, policyProfile := range policyProfiles {
			data[i] = policyProfile
		}
		format.PrettyPrintTable(data,
			"id",
			"name",
			"organizationName",
			"forbidHttpIngress",
			"allowedRepo",
			"forbidNodePort",
			"forbidSpecificTags",
			"ingressWhitelist",
			"requireProbe",
			"uniqueIngresses",
			"uniqueServiceSelector",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := opa_profiles.NewOpaProfilesListParams().WithV(apiconfig.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		apiconfig.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&apiconfig.SortDirection)
	}

	var policyProfiles = make([]*models.OpaProfileListDto, 0)
	for {
		response, err := apiClient.Client.OpaProfiles.OpaProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		policyProfiles = append(policyProfiles, response.Payload.Data...)
		count := int32(len(policyProfiles))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(policyProfiles)) > opts.Limit {
		policyProfiles = policyProfiles[:opts.Limit]
	}

	printResults(policyProfiles)
	return
}
