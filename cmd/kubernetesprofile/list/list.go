package list

import (
	"github.com/itera-io/taikun-cli/api"
	"github.com/itera-io/taikun-cli/cmd/cmdutils"
	"github.com/itera-io/taikun-cli/config"
	"github.com/itera-io/taikun-cli/utils/out"

	"github.com/itera-io/taikungoclient/client/kubernetes_profiles"
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
		Short: "List kubernetes profiles",
		RunE: func(cmd *cobra.Command, args []string) error {
			return listRun(&opts)
		},
		Args:    cobra.NoArgs,
		Aliases: cmdutils.ListAliases,
	}

	cmd.Flags().Int32VarP(&opts.OrganizationID, "organization-id", "o", 0, "Organization ID (only applies for Partner role)")

	cmdutils.AddLimitFlag(cmd)
	cmdutils.AddSortByAndReverseFlags(cmd, models.KubernetesProfilesListDto{})

	return cmd
}

func listRun(opts *ListOptions) (err error) {
	apiClient, err := api.NewClient()
	if err != nil {
		return
	}

	params := kubernetes_profiles.NewKubernetesProfilesListParams().WithV(api.Version)
	if opts.OrganizationID != 0 {
		params = params.WithOrganizationID(&opts.OrganizationID)
	}
	if config.SortBy != "" {
		params = params.WithSortBy(&config.SortBy).WithSortDirection(api.GetSortDirection())
	}

	var kubernetesProfiles = make([]*models.KubernetesProfilesListDto, 0)
	for {
		response, err := apiClient.Client.KubernetesProfiles.KubernetesProfilesList(params, apiClient)
		if err != nil {
			return err
		}
		kubernetesProfiles = append(kubernetesProfiles, response.Payload.Data...)
		count := int32(len(kubernetesProfiles))
		if config.Limit != 0 && count >= config.Limit {
			break
		}
		if count == response.Payload.TotalCount {
			break
		}
		params = params.WithOffset(&count)
	}

	if config.Limit != 0 && int32(len(kubernetesProfiles)) > config.Limit {
		kubernetesProfiles = kubernetesProfiles[:config.Limit]
	}

	out.PrintResults(kubernetesProfiles,
		"id",
		"name",
		"organizationName",
		"taikunLBEnabled",
		"octaviaEnabled",
		"exposeNodePortOnBastion",
		"cni",
		"allowSchedulingOnMaster",
	)
	return
}
