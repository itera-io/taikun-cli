package list

import (
	"fmt"

	"taikun-cli/api"
	"taikun-cli/config"
	"taikun-cli/utils"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
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
		Short: "List kubernetes profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			if opts.Limit < 0 {
				return utils.NegativeLimitFlagError
			}
			if !config.OutputFormatIsValid() {
				return config.OutputFormatInvalidError
			}
			return listRun(&opts)
		},
		Args: cobra.NoArgs,
	}

	cmd.Flags().BoolVarP(&opts.ReverseSortDirection, "reverse", "r", false, "Reverse order of results")
	cmd.Flags().Int32VarP(&opts.Limit, "limit", "l", 0, "Limit number of results (limitless by default)")
	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")
	cmd.Flags().StringVarP(&opts.SortBy, "sort-by", "s", "", "Sort results by attribute value")

	return cmd
}

func printResults(kubernetesProfiles []*models.KubernetesProfilesListDto) {
	if config.OutputFormat == config.OutputFormatJson {
		utils.PrettyPrintJson(kubernetesProfiles)
	} else if config.OutputFormat == config.OutputFormatTable {
		data := make([]interface{}, len(kubernetesProfiles))
		for i, kubernetesProfile := range kubernetesProfiles {
			data[i] = kubernetesProfile
		}
		utils.PrettyPrintTable(data,
			"id",
			"name",
			"organizationName",
			"taikunLBEnabled",
			"octaviaEnabled",
			"exposeNodePortOnBastion",
			"cni",
			"allowSchedulingOnMaster",
		)
	}
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kubernetes_profiles.NewKubernetesProfilesListParams().WithV(utils.ApiVersion)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if opts.ReverseSortDirection {
		utils.ReverseSortDirection()
	}
	if opts.SortBy != "" {
		params = params.WithSortBy(&opts.SortBy).WithSortDirection(&utils.SortDirection)
		fmt.Printf("sorting by %s\n", opts.SortBy)
	}

	var kubernetesProfiles = make([]*models.KubernetesProfilesListDto, 0)
	for {
		response, err := apiClient.Client.KubernetesProfiles.KubernetesProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		kubernetesProfiles = append(kubernetesProfiles, response.Payload.Data...)
		count := int32(len(kubernetesProfiles))
		if opts.Limit != 0 && count >= opts.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if opts.Limit != 0 && int32(len(kubernetesProfiles)) > opts.Limit {
		kubernetesProfiles = kubernetesProfiles[:opts.Limit]
	}

	printResults(kubernetesProfiles)
	return
}
